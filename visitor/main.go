package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
    VisitDoubleExpression(e *DoubleExpression)
    VisitAdditionExpression(e *AdditionExpression)
}
type Expression interface {
    Accept(ev ExpressionVisitor)
}
type DoubleExpression struct {
    value float64
}
type AdditionExpression struct {
 left, right Expression
}
func (d *DoubleExpression) Print(sb *strings.Builder) {
    sb.WriteString(fmt.Sprintf("%g", d.value))
}
// func (a *AdditionExpression) Print(sb *strings.Builder) {
//     sb.WriteRune('(')
//     a.left.Print(sb)
//     sb.WriteRune('+')
//     a.right.Print(sb)
//     sb.WriteRune(')')
// }
func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
    ev.VisitDoubleExpression(d)
}
func (d *AdditionExpression) Accept(ev ExpressionVisitor) {
    ev.VisitAdditionExpression(d)
}
type ExpressionPrinter struct {
    sb strings.Builder
}
func (e *ExpressionPrinter) VisitDoubleExpression(d *DoubleExpression) {
    e.sb.WriteString(fmt.Sprintf("%g", d.value))
}
func (e *ExpressionPrinter) VisitAdditionExpression(a *AdditionExpression) {
    e.sb.WriteRune('(')
    a.left.Accept(e)
    e.sb.WriteRune('+')
    a.right.Accept(e)
    e.sb.WriteRune(')')
}
func NewExpressionPrinter() *ExpressionPrinter {
    return &ExpressionPrinter{strings.Builder{}}
}
func (ep *ExpressionPrinter) String() string {
    return ep.sb.String()
}
type ExpressionEvaluator struct {
    result float64
}
func (e *ExpressionEvaluator) VisitDoubleExpression(d *DoubleExpression) {
    e.result = d.value
}
func (e *ExpressionEvaluator) VisitAdditionExpression(a *AdditionExpression) {
   a.left.Accept(e)
   x := e.result
   a.right.Accept(e)
   x += e.result
   e.result = x
}


func main() {
    // 1 + (2+3)
    e := AdditionExpression{
        left: &DoubleExpression{1},
        right: &AdditionExpression{
            left: &DoubleExpression{2},
            right: &DoubleExpression{3},
        },
    }
    ep := NewExpressionPrinter()
    e.Accept(ep)
    ee := &ExpressionEvaluator{}
    e.Accept(ee)
    fmt.Printf("%s = %g", ep, ee.result)

}
