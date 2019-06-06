# go-veevalidate

Create Veevalidate rules in go.

Example:

````

v := veevalidate.New().
		Required().
		MaxValue(100).
		MaxValue(0xffff),
	Default: 502,
}

````
