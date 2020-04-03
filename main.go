package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)
type player struct {
	image *ebiten.Image
	xPos,yPos float64
	speed float64
}
var(
	err error
	background *ebiten.Image
	spaceShip *ebiten.Image
	playerOne *player
)

const (
	screenWidth,screenHeight = 640,400
)
func init(){
	background,_,err = ebitenutil.NewImageFromFile("assets/game-background-png-6.png",ebiten.FilterDefault)
	if err!=nil{
		log.Fatal(err)
	}
	spaceShip,_,err = ebitenutil.NewImageFromFile("assets/rocket32.png",ebiten.FilterDefault)
	if err!=nil{
		log.Fatal(err)
	}
	playerOne = &player{spaceShip,screenWidth/2,screenHeight/2,4}
}
func movePlayer(){
	if ebiten.IsKeyPressed(ebiten.KeyUp){
		playerOne.yPos-=playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown){
		playerOne.yPos+=playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		playerOne.xPos-=playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
		playerOne.xPos+=playerOne.speed
	}
}
func update(screen *ebiten.Image) error{
	movePlayer()
	if ebiten.IsDrawingSkipped(){
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0,0)
	err = screen.DrawImage(background,op)
	if err!=nil{
		return err
	}
	playerOp:=&ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos,playerOne.yPos)
	screen.DrawImage(playerOne.image,playerOp)
	return nil

}
func main(){
	if err:=ebiten.Run(update,screenWidth,screenHeight,2,"Game Dev");err!=nil{
		log.Fatal(err)
	}
}