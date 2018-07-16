package main

import (
	"fmt"
	"flag"
	"context"
	"strconv"
	"net/http"
	"io/ioutil"
	"io"
	"os"
	"html/template"
	"log"

	"github.com/labstack/echo"



	tf_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	"google.golang.org/grpc"

	"path/filepath"
)

var (
	port         int
	model_server string
	model_name   string
)

func init() {
	flag.IntVar(&port, "port", 1323, "concurrent processing ,default 1 .")
	flag.StringVar(&model_server, "model_server", "model-server:8500", "concurrent processing ,default 1 .")
	//flag.StringVar(&model_name, "model_name", "default", "concurrent processing ,default 1 .")
	flag.StringVar(&model_name, "model_name", "inception", "concurrent processing ,default 1 .")
	flag.Parse()
}

type Resp struct {
	Success bool
	Msg     string
	Result  [4]float32
}

var buf []byte

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.gptl")
		checkErr(err)
		t.Execute(w, nil)
	} else {
		file, handle, err := r.FormFile("file")
		checkErr(err)
		f, err := os.OpenFile("./pic/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)
		checkErr(err)
		defer f.Close()
		defer file.Close()
		fmt.Println("upload success!")
	}
}

func checkErr(err error) {
	if err != nil {
		err.Error()
	}
}

func main() {

	e := echo.New()
	// Set up a connection to the model server.
	conn, err := grpc.Dial(model_server, grpc.WithInsecure())
	if err != nil {
		e.Logger.Fatalf("can not connect model_server: %v", conn)
	}
	defer conn.Close()

	client := pb.NewPredictionServiceClient(conn)

	//homedir, err := homedir.Dir()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//print(homedir)
	print(dir)

	e.Static("/", dir+"/templates")
	e.Static("/pic",dir+"/pic")
	e.Static("/webuploader", dir+"/static/webuploader")
	e.Static("/static", dir+"/static/mnist")

	e.POST("/api/mnist", func(c echo.Context) error {
		req := c.Request()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		//result, err := Predict(client, body)
		result, err := Predict2(client, body)

		if err != nil {
			e.Logger.Error(err.Error())
			return c.JSON(http.StatusOK, &Resp{
				Msg: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, &Resp{
			Success: true,
			Result:  result,
		})
	})
	e.POST("/api/upload", func(c echo.Context) error {
		req := c.Request()
		req.ParseForm()

		file, handle, err := req.FormFile("file")
		checkErr(err)
		//f, err := os.OpenFile("./test/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(dir+"/pic/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)
		checkErr(err)
		defer f.Close()
		defer file.Close()
		fmt.Println("upload success")

		return c.JSON(http.StatusOK, &Resp{
			Success: true,
			Result:  [4]float32{0.0, 0.0, 0.0, 0.0},
		})
	})

	e.POST("/api/flower", func(c echo.Context) error {

		req := c.Request()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		print("."+string(body))

		imageBytes, err := ioutil.ReadFile(dir+string(body))
		result, err := Predict2(client, imageBytes)

		if err != nil {
			e.Logger.Error(err.Error())
			return c.JSON(http.StatusOK, &Resp{
				Msg: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, &Resp{
			Success: true,
			Result:  result,
		})
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

func Predict2(c pb.PredictionServiceClient, imgBytes []byte) (result [4]float32, err error) {
	req := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{Name: model_name, SignatureName: "predict"},
		Inputs:    make(map[string]*tf_framework.TensorProto),
	}
	imageTensor, err := tf.NewTensor(string(imgBytes))
	if err != nil {
		log.Fatalln("Cannot read image file")
	}
	tensorString, ok := imageTensor.Value().(string)
	if !ok {
		log.Fatalln("Cannot type assert tensor value to string")
	}

	tp := &tf_framework.TensorProto{
		Dtype: tf_framework.DataType_DT_STRING,
		TensorShape: &tf_framework.TensorShapeProto{

		},
		StringVal: [][]byte{[]byte(tensorString)},
	}
	req.Inputs["inputs"] = tp

	resp, err := c.Predict(context.Background(), req)
	if err != nil {
		return
	}
	output, ok := resp.Outputs["pred"]
	if !ok {
		err = fmt.Errorf("can not find output data with label pred")
		return
	}
	if len(output.FloatVal) != 4 {
		err = fmt.Errorf("wrong output dimension, it should be 5, now got %d", len(output.FloatVal))
		return
	}
	copy(result[:], output.FloatVal)
	return
}

func Predict(c pb.PredictionServiceClient, imgBytes []byte) (result [10]float32, err error) {
	req := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{Name: model_name},
		Inputs:    make(map[string]*tf_framework.TensorProto),
	}
	in := normalize(imgBytes)

	tp := &tf_framework.TensorProto{
		Dtype:    tf_framework.DataType_DT_FLOAT,
		FloatVal: in,
		TensorShape: &tf_framework.TensorShapeProto{
			Dim: []*tf_framework.TensorShapeProto_Dim{
				&tf_framework.TensorShapeProto_Dim{
					Size: int64(1),
					Name: "batch",
				},
				&tf_framework.TensorShapeProto_Dim{
					Size: int64(28),
					Name: "x",
				},
				&tf_framework.TensorShapeProto_Dim{
					Size: int64(28),
					Name: "y",
				},
				&tf_framework.TensorShapeProto_Dim{
					Size: int64(1),
					Name: "channel",
				},

			},
		},
	}
	req.Inputs["x"] = tp

	resp, err := c.Predict(context.Background(), req)
	if err != nil {
		return
	}
	output, ok := resp.Outputs["y"]
	if !ok {
		err = fmt.Errorf("can not find output data with label y")
		return
	}
	if len(output.FloatVal) != 10 {
		err = fmt.Errorf("wrong output dimension, it should be 10, now got %d", len(output.FloatVal))
		return
	}
	copy(result[:], output.FloatVal)
	return
}

func normalize(bytes []byte) (r []float32) {
	r = make([]float32, 0, len(bytes))
	for _, b := range bytes {
		r = append(r, float32(255-b)/255)
	}
	return
}

func normalize2(bytes []byte) (r []float32) {
	r = make([]float32, 0, len(bytes))
	for _, b := range bytes {
		r = append(r, float32(b-128)/128)
	}
	print(len(r))
	return
}