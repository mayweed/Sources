#!/usr/bin/python

import tkinter
import xml
import sys
import turtle

#class needed for graphicsCommands
class PyList:
    def __init__(self):
        self.items=[]
    def append(self,item):
        self.items=self.items + [item]

    #THIS piece of code very interesting!!
    def __iter__(self):
        for c in self.items:
            yield c


class DrawingApplication(tkinter.Frame):

    def __init__(self, master=None):
        super().__init__(master)
        self.pack()
        self.buildWindow()
        self.graphicsCommands=PyList()
    
    def buildWindow(self):
        self.master.title("Draw")
        bar=tkinter.Menu(self.master)
        # tearoff: cant be separated from window
        fileMenu=tkinter.Menu(bar,tearoff=0)

        #to avoid the "UnboundLocalVariable" error
        canvas=tkinter.Canvas(self,width=600,height=600)
        canvas.pack(side=tkinter.LEFT)

        theTurtle=turtle.RawTurtle(canvas)
        theTurtle.shape("circle")
        screen=theTurtle.getscreen()
        
        # Do not update the screen unless update() is called
        screen.tracer(0)

        def NewWindow():
            theTurtle.clear()
            theTurtle.penup()
            theTurtle.goto(0,0)
            theTurtle.pendown()
            screen.update()
            screen.listen()
            self.graphicsCommands=PyList()
        # link a function to a Menu command!!
        fileMenu.add_command(label="New",command=NewWindow())

        def parse(filename):
            xmldoc=xml.dom.minidom.parse(filename)
            graphicsCommandsElement=xmldoc.getElementsByTagName("GraphicsCommands")[0]
            graphicsCommands=graphicsCommandsElement.getElementByTagName("Command")

            for commandElement in graphicsCommands:
                print(type(commandElement))
                command=commandElement.firstChild.data.strip()
                attr=commandElement.attributes
                if command=="GoTo":
                    x=float(attr["x"].value)
                    y=float(attr["y"].value)
                    width=float(attr["width"].value)
                    color=attr["color"].value.strip()
                    cmd=GoToCommand(x,y,width,color)
                elif command=="Circle":
                    radius=float(attr["radius"].value)
                    width=float(attr["width"].value)
                    color=attr["color"].value.strip()
                    cmd=CircleCommand(radius,width,color)
                elif command=="BeginFill":
                    color=attr["color"].value.strip()
                    cmd=BeginFill(color)
                elif command=="EndFill":
                    cmd=EndFillCommand()
                elif command == "PenUp":
                    cmd=PenUpCommand()
                elif command=="PenDown":
                    cmd=PenDownCommand()
                else:
                    raise RuntimeError("Unknown Command: "+ command)
                self.graphicsCommands.append(cmd)

        def loadFile():
            filename=tkinter.filedialog.askopenfilename(title="Select a Graphics File")
            newWindow()
            self.graphicsCommands=PyList()
            parse(filenmae)

            for cmd in self.graphicsCommands:
                cmd.draw(theTurtle)

            screen.update()

        fileMenu.add_command(label="Load...",command=loadFile)

        def addToFile():
            filename=tkinter.filedialog.askopenfilename(title="Select a Graphics File")
            theTurtle.penup()
            theTurtle.goto(0,0)
            theTurtle.pendown()
            theTurtle.pencolor("#000000")
            theTurtle.fillcolor("#000000")
            cmd=PenUpCommand()
            self.graphicsCommands.append(cmd)
            cmd=GoToCommand(0,0,1,"#000000")
            self.graphicsCommands.append(cmd)
            cmd=PenDownCommand()
            self.graphicsCommands.append(cmd)
            screen.update()
            parse(filename)

            for cmd in self.graphicsCommands:
                cmd.draw(theTurtle)

            screen.update
        fileMenu.add_command(label="Load into...",command=addToFile)

        def write(filename):
            file=open(filename,"w")
            file.write('<?xml version="1.0" encoding="UTF-8" standalone="no" ?>\n')
            file.write('<GraphicsCommands>\n')
            for cmd in self.graphicsCommands:
                file.write('    '+str(cmd)+"\n")
            file.write('</GraphicsCommands>\n')
            file.close

        def saveFile():
            filename=tkinter.filedialog.asksaveasfilename(title="Save Picture As...")
            write(filename)

        fileMenu.add_command(label="Exit",command=self.master.quit)
        bar.add_cascade(label="File",menu=fileMenu)
        self.master.config(menu=bar)

        #canvas=tkinter.Canvas(self,width=600,height=600)
        #canvas.pack(side=tkinter.LEFT)

        #theTurtle=turtle.RawTurtle(canvas)
        #theTurtle.shape("circle")
        #screen=theTurtle.getscreen()
        
        # Do not update the screen unless update() is called
        #screen.tracer(0)
        
        sideBar=tkinter.Frame(self,padx=5,pady=5)
        sideBar.pack(side=tkinter.RIGHT,fill=tkinter.BOTH)
        pointLabel=tkinter.Label(sideBar,text="Width")
        pointLabel.pack()

        widthSize=tkinter.StringVar()
        widthEntry=tkinter.Entry(sideBar,textvariable=widthSize)
        widthEntry.pack()
        widthSize.set(str(1))

        radiusLabel=tkinter.Label(sideBar,text="Radius")
        radiusLabel.pack()
        radiusSize=tkinter.StringVar()
        radiusEntry=tkinter.Entry(sideBar,textvariable=radiusSize)
        radiusSize.set(str(10))
        radiusEntry.pack()

        def circleHandler():
            cmd=CircleCommand(float(radiusSize.get()),float(widthSize.get()),penColor.get())
            cmd.draw(theTurtle)
            self.graphicsCommands.append(cmd)

            screen.update()
            screen.listen()

        circleButton=tkinter.Button(sideBar,text="Draw Circle", command=circleHandler)
        circleButton.pack(fill=tkinter.BOTH)

        screen.colormode(255)
        penLabel=tkinter.Label(sideBar,text="Pen Color")
        penLabel.pack()
        penColor=tkinter.StringVar()
        penEntry=tkinter.Entry(sideBar,textvariable=penColor)
        penEntry.pack()
        penColor.set("#000000")

        def getPenColor():
            color=tkinter.colorchooser.askcolor()
            if color!=None:
                penColor.set(str(color)[-9:-2])

        penColorButton=tkinter.Button(sideBar,text="Pick Pen Color", command=getPenColor)
        penColorButton.pack(fill=tkinter.BOTH)

        fillLabel=tkinter.Label(sideBar,text="Fill Color")
        fillLabel.pack()
        fillColor=tkinter.StringVar()
        fillEntry=tkinter.Entry(sideBar,textvariable=fillColor)
        fillEntry.pack()
        fillColor.set("#000000")

        def getFillColor():
            color=tkinter.colorchooser.askcolor()
            if color != None:
                fillColor.set(str(color)[-9:-2])

        fillColorButton=tkinter.Button(sideBar,text="Pick Fill Color",command=getFillColor)
        fillColorButton.pack(fill=tkinter.BOTH)

        def beginFillHandler():
            cmd=BeginFillCommand(fillColor.get())
            cmd.draw(theTurtle)
            self.graphicsCommands.append(cmd)

        beginFillButton=tkinter.Button(sideBar,text="Begin Fill",command=beginFillHandler)
        beginFillButton.pack(fill=tkinter.BOTH)

        def endFillHandler():
            cmd=EndFillCommand()
            cmd.draw(theTurtle)
            self.graphicsCommands.append(cmd)

        endFillButton=tkinter.Button(sideBar,text="End Fill",command=endFillHandler)
        endFillButton.pack(fill=tkinter.BOTH)

        penLabel=tkinter.Label(sideBar,text="Pen is Down")
        penLabel.pack()

        def penUpHandler():
            cmd=PenUpCommand()
            cmd.draw(theTurtle)
            penLabel.configure(text="Pen is Up")
            self.graphicsCommands.append(cmd)

        penUpButton=tkinter.Button(sideBar,text="Pen Up",command=penUpHandler)
        penUpButton.pack(fill=tkinter.BOTH)

        def penDownHandler():
            cmd=PenDownCommand()
            cmd.draw(theTurtle)
            penLabel.configure(text="Pen is Down")
            self.graphicsCommands.append(cmd)

        penDownButton=tkinter.Button(sideBar,text="Pen Down",command=penDownHandler)
        penDownButton.pack(fill=tkinter.BOTH)

        def clickHandler(x,y):
            cmd=GoToCommand(x,y,float(widthSize.get()),penColor.get())
            cmd.draw(theTurtle)
            self.graphicsCommands.append(cmd)
            screen.update()
            screen.listen()

        screen.onclick(clickHandler)

        def dragHandler(x,y):
            cmd=GoToCommand(x,y,float(widthSize.get()),penColor.get())
            cmd.draw(theTurtle)
            self.graphicsCommands.append(cmd)
            screen.update()
            screen.listen()

        theTurtle.ondrag(dragHandler)

        def undoHandler():
            if len(self.graphicsCommands) > 0:
                self.graphicsCommands.removeLast()
                theTurtle.clear()
                theTurtle.penup()
                theTurtle.goto(0,0)
                theTurtle.pendown()
                for cmd in self.graphicsCommands:
                    cmd.draw(theTurtle)
                screen.update()
                screen.listen()

            screen.onkeypress(undohandler,"u")
            screen.listen()

def main():
    root=tkinter.Tk()
    drawingApp=DrawingApplication(root)
    
    drawingApp.mainloop()
    print("Program Execution Completed.")

if __name__ == "__main__":
    main()

