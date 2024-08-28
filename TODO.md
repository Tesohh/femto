# Important stuff

- [ ] Operations on line (yy, cc, dd, ...) in normal mode
- [ ] Visual mode
- [ ] Operations (y, c, d, ...) in visual mode
- [ ] Range based drawing

# Last stuff

- [ ] add all the fun features

# Secondary stuff

- [ ] Config (loading from file) (do we even really need a config??)
- [ ] Plugin viewer
- [ ] Composite motions (eg c then iw, c then w, d then d)
  - [ ] Could have something where when c is pressed it starts catching events like the hidden "o" mode in vim, and when it's satisfied exceute command with motion and go back into normal mode
  - [ ] Probably scope creep. Please just add a visual mode with operations and thats it

# Backlog

- [x] Move from ebiten to terminal rendering
- [x] Human key package
- [x] Basic editing features
  - [x] hjkl and arrow keys
  - [x] Mode switching
  - [x] insert mode
  - [x] file writing / reading
- [ ] Windowing system
  - [x] Orchestrating
  - [ ] Rendering
    - [x] basic drawing
    - [ ] range based drawing (for when content goes out of rendering bounds)
    - [x] render cursor based on current window start
      - [x] What if the cursor is managed by the window that has focus? we can pass a isFocused parameter...
  - [x] Move to editor being a window too
    - [x] Center window
    - [x] The buffer is a tab local window
    - [x] Movement commands are easy: they should apply to ANY buffer, just make Buf() return the currently focused buffer
    - [x] Editing commands too should simply apply to the current buffer (just that the buffer can be written to)
    - [x] Where do we store filenames and other editor's tab's buffer specific data?
      - [x] Also easy, we move it from the Tab to the Window.
      - [x] UI buffers have no filename, which means they are "scratchpads"
      - [x] Mode, sequence and all that is saved in the Window of course.
      - [x] Tab becomes just a collection of tab local windows
  - [x] Ability to get a window by id, both on editor and tab
- [x] Plugin "api"
  - [x] Base Plugin interface
  - [x] DumbPlugin implementation
  - [x] Access to editor stuff
- [x] interactive windows
  - [x] Each window has it's own humankeymap and commandmap, which:
    - [x] the keymap gets merged on the global one
    - [x] when using e.RunCommand, check for window commands and give priority to those
- [x] Color sections
- [x] Theming
- [x] Command line
- [x] Error logging on cmdbar
- [x] Commands with arguments

- [x] Fix: pointers returned by RegisterWindow act weird.
- [x] Fix: GetWindow also has weird pointers, probably due to it returning a pointer to the for loop's `w`
- [x] Fix: commandbar commands get merged PERMANENTLY into the editor's commandmap it seems.
      When you type a command and execute it, then press ESC in insert mode, it tries to do command.exit
