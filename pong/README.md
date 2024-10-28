
# Setup

```shell
# Get sdl2 lib
sudo apt-get install libsdl2-dev

go mod tidy

go build -o pong pong.go

./pong

```

# ToDo

Features

- [x] logical order code into separate files 
- [ ] Game over state
- [x] paddle stay on screen
- [ ] ball cannot bounce behind paddle..
- [x] start state: paddle in the middle

- [x] Build app, make it distributable
- [ ] Package the window/draw/ball/rectangle utils

- [ ] AI needs to be inperfect
- [ ] improve gameplay
    - paddle angle bouncing
    - Ball start in random direction
    - [x] Ball velocity increases
- [x] two player or PC
- [ ] resizing of the window
