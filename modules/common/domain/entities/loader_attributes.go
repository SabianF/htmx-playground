package common

import "reflect"

type LoaderAttributes struct {
  AriaLabel string
  Color string
  Delay string
  Height string
  Width string
}

func (la *LoaderAttributes) ColorName() string {
  return reflect.TypeOf(la.Color).Elem().Name()
}

func (la *LoaderAttributes) HeightName() string {
  return reflect.TypeOf(la.Height).Elem().Name()
}

func (la *LoaderAttributes) WidthName() string {
  return reflect.TypeOf(la.Width).Elem().Name()
}

func (la *LoaderAttributes) HasAriaLabel() bool {
  return len(la.AriaLabel) > 0
}

func (la *LoaderAttributes) HasColor() bool {
  return len(la.Color) > 0
}

func (la *LoaderAttributes) HasDelay() bool {
  return len(la.Delay) > 0
}

func (la *LoaderAttributes) HasHeight() bool {
  return len(la.Height) > 0
}

func (la *LoaderAttributes) HasWidth() bool {
  return len(la.Width) > 0
}

func (la *LoaderAttributes) HasCSS() bool {

  if la.HasColor() {
    return true
  }

  if la.HasHeight() {
    return true
  }

  if la.HasWidth() {
    return true
  }

  return false
}

func (la *LoaderAttributes) GetCSS() string {

  var css string = ""

  if la.HasColor() {
    css += "background-color:" + la.Color + ";"
  }

  if la.HasHeight() {
    css += "height:" + la.Height + ";"
  }

  if la.HasWidth() {
    css += "width:" + la.Width + ";"
  }

  return css
}
