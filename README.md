# golang-cube
im learning math!



## 3D cube rendering (wireframe) using RPC and multiple services

Interface module takes care of graphics + kb input and reports key up/down events to the math engine

Math engine converts the points + orientation/rotation to draw commands (lines polygons and points) 


GUI captures key commands -> math module computes 3D it -> render module turns it to 2D points and lines -> GUI module displays it
