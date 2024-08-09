# Femto

Worse than nano!

![femto logo](assets/appicon.png)

---

Esoteric editor along the lines of [vim](https://vimto.s3.eu-west-1.amazonaws.com/wp-content/uploads/2021/03/31111534/Middle-East-v5.png) and Fortnite.

It's so bad, you're just better off writing on paper and OCR scanning your code

# Features

- Actually good features:
  - Fully modular:
    - The editor is completely based on a plugin system.
    - Almost everything, (except the very core of the editor main loop) is a plugin
      - (even the insert mode is a plugin)
    - Plugins that don't need to be injected in the main loop, and just define commands and keymaps can use the DumbPlugin struct.
    - More complex plugins crate a struct that implements the Plugin interface
      - This means that plugins can even be stateful by simply having properties in the struct.
    - Plugins can communicate with eachother through events
  - Portable:
    - femto compiles into a single executable, including plugins and library dependencies.
      - caveat: you need to recompile to add plugins.
  - Modal
  - Windowing system:
    - Windows are placed automatically by the editor and can be managed by plugins very easily
      - calling editor.RegisterWindow returns a pointer to the registered window so plugins can store the window and use it in their Update or Draw functions
    - The text editor itself is just a writeable window
    - Each window can add their own, window-local Commands and Keymaps.
      - Imagine for example a netrw like window with it's own keymaps.
      - You can also disable a key here, by mapping it to the `noop` command (eg for disabling insert mode)
- Esoteric features (planned):
  - Optional plugins:
    - [ ] Battle pass: locks some editor features behind a leveling system, the ideal way to use a text editor.
    - [ ] Goals: Can set daily goals, like "use femto for 30 min" and gain battle pass xp from it.
    - [ ] Advertising: provides passive battle pass xp through advertisements in several toolbars
    - [ ] Starfieldify: locks the editor to 1 fps.
    - [ ] AIGPTSLOP9K: uses AI™ machine™ learning™ llm™ predictive™ interactive™ technology™ to suggest your next character™. your keyboard will have 1 key™ and you will like it!
    - [ ] NFT: gives you an NFT every battle pass level.
    - [ ] Quantum: uses quantum 999999 quettaflops of cloud quantum computing power to greet users when entering the editor.
    - [ ] CrowdStrikeStrike: There's a very small chance that your editor goes into a BSOD loop, unless you restart it 15 times.
    - [ ] JSGod: Asks the [javascript god](https://www.youtube.com/watch?v=6FEYcBPBGOk&t=428s) how much time passed since his latest creation (framework) when you open a javascript/typescript file.
    - [ ] Electronify: your editor will load 4s slower and use 400mb more RAM.
