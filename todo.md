# To Do

## Next (09/27/2023)

- Fonts
  - Using nerdfont glyphs, but debugged just uses natural values
  - Add attributes & flag them
- Multiple Layers
  - Active layer
  - Collapsing layer into parent
  - Removing layer

## Older

- Pause the program via not reading the FIFO when the program writes to it.
  - So, could implement pausing on every frame, for example.
- Put all attributes on the Cell.
  - bg, attrs (underline, dim, italics, etc.)
- Figure out the Unicode ids for things like the arrows.
- termgridboard (tgb) monitors /tmp/termgridboard for new file(s). When they show up,
  tgb then opens the FIFO for reading.
  - So, the program, by creating the FIFO indicates the desire to be monitored.

## TODO in Truss

- Right now, we are just sending the MxN grid values each 'frame'.
  - Should be sending the initial state of the grid in the beginning, but
    then only sending diffs when events happen in the logic.
- Should send which cells are 'examined', when the robot looks in a direction.




