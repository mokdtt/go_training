package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go_training/ch07/ex13/eval"
)

func parse(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func checkAndEval(expr eval.Expr, r *http.Request) (float64, error) {
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return 0, err
	}
	varsValue := eval.Env{}
	for v := range vars {
		valueStr := r.Form.Get(string(v))
		if valueStr == "" {
			return 0, fmt.Errorf("variable %s not found", v)
		}
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return 0, err
		}
		varsValue[v] = value
	}
	return expr.Eval(varsValue), nil
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parse(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	res, err := checkAndEval(expr, r)
	if err != nil {
		http.Error(w, "bad variable: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s = %g\n", expr.String(), res)
	vars := make(map[eval.Var]bool)
	expr.Check(vars)
	for v := range vars {
		fmt.Fprintf(w, "%s = %s\n", v, r.Form.Get(string(v)))
	}
}

func main() {
	http.HandleFunc("/calc", calc)
	fmt.Println("example: http://localhost:8000/calc?expr=pow(a,b)&a=2&b=3")
	fmt.Println("メモ: このままだとプラスが受け取れんかも？")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
