# Sherlock - Google Wire Dependency Injection Struct Auto Discovery


## Install
```bash
go get github.com/alfiankan/sherlock-struct-autowire
```

## Comparison

### Manual Dependency Injection
```go
func InitializeCar(pertalite example.Pertalite, body example.Body) example.Car {
	gas := example.Gas{
		GasOil: pertalite,
	}
	oil := example.Oil{}
	engine := example.Engine{
		Gas: gas,
		Oil: oil,
	}
	car := example.Car{
		Engine: engine,
		Body:   body,
	}
	return car
}

func InitializeEngine(int64_2 int64) example.Engine {
	pertalite := example.Pertalite{
		Barrel: int64_2,
	}
	gas := example.Gas{
		GasOil: pertalite,
	}
	oil := example.Oil{}
	engine := example.Engine{
		Gas: gas,
		Oil: oil,
	}
	return engine
}

func InitializeBody() example.Body {
	body := example.Body{}
	return body
}
```

### Google Wire
```go
func InitializeCar(pertalite example.Pertalite, body example.Body) example.Car {
	wire.Build(
		wire.Struct(new(example.Car), "*"),
		wire.Struct(new(example.Engine), "*"),
		wire.Struct(new(example.Gas), "*"),
		wire.Struct(new(example.Oil), "*"),
	)
	return example.Car{}
}

func InitializeEngine(int64 int64) example.Engine {
	wire.Build(
		wire.Struct(new(example.Engine), "*"),
		wire.Struct(new(example.Gas), "*"),
		wire.Struct(new(example.Pertalite), "*"),
		wire.Struct(new(example.Oil), "*"),
	)
	return example.Engine{}
}
```

### Google Wire with sherlock
```go
	sherlock.New().
		SetPkgName("tests").
		Add(example.Car{}, customBody).
		Add(example.Engine{}).
		Add(example.Body{}).
		SetGlobalInject(config).
		Gen()
```

## Create Dependency

```go
	sherlock.New(). // new sherlock
		SetPkgName("tests"). // set pkg name, is main by default
		Add(example.Car{}, customBody). // add auto wiring deps
		Add(example.Engine{}). // add auto wiring deps
		Add(example.Body{}). // add auto wiring deps
		SetGlobalInject(config). // set global overrider, injector, provider
		Gen()


```

## Create Simple Deps

```go
    sherlock.New().Add(example.Car{})
```

sherlock willinspect required struct field recursively

## Create Simple Deps with override

```go
    sherlock.New().Add(example.Car{})
```

sherlock willinspect required struct field recursively, in case we need to override some dependencies jut add variadic object next to first deps.

```go
    pertalite := Pertalite{Barel: 5000}
    sherlock.New().Add(example.Car{}, pertalite)
```

## Create Simple Deps with global override

```go
    pertalite := Pertalite{Barel: 5000}
    sherlock.New().
        Add(example.Car{}).
        Add(example.Engine{}).
        SetGlobalInject(pertalite).
```

Global inject will inject to all dependecies



## How To Install
- Requirements
    - Google Wire https://github.com/google/wire
    - Go 1.18+

- Install Sherlock 
    ```
    go get ....
    ```


## Example
