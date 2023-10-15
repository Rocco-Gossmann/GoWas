Bugs:
================================================================================

TextDisplay:
================================================================================
    [x] Add a clear function, that replaces a number of characters with spaces, but leaves the Cursor-Position unaltered
    [ ] Make Text-Display Moveable on Screen Similar to TileMap [ ] Add Setters and Getters like in BitmapEntity
        [ ] ScreenOffset 
        [ ] Alpha

Rendering:
================================================================================
    [x] Test how drawing a bitmap, that overflows all 4 sides functions @critical 
        [x] Fix Oversized Bitmaps being cut of to early
    [ ] Adapt TileSets for SpriteSheets
    [ ] Tinting Bitmaps/TileMaps/Sprites on Draw


Maps:
================================================================================
    [ ] Rework Setters and Getters to Work like BitmapEntities Do @critical
    [ ] Test how drawing Maps, that are smaller than the screen behaves @critical
    [ ] Maps Collision layers @critical
    [ ] Maps Alpha blending
    [ ] Allow to fill columns or rows of Tiles with data @high (Instead of just single tiles or the entire map)
    [ ] Maps Clipping-Rect @low
    [ ] Allow maps scroll.x any y propperties to be negative @low


Keyboard Input:
================================================================================
    [ ] Browser side - Key-Statemanagement
    [ ] Go Side
        [ ] KeyStates


Audio:
================================================================================
    [ ] Tell Browser to load and provide Audio from within Go <audio> tag
        [ ] Establish some sort of handle that can be used to identify the sound in both JS and Go
    [ ] Trigger audio playback from within Go
    [ ] Cancel audio playback from within Go
    [ ] Mabe change audio volumne from within Go @low 


Tooling:
================================================================================
    [x] Automate Assets.png cnnversion (User should only have to put the PNG into the ./assets folder)
        [x] Extend Makefile
        [x] Move tools-scripts into "project-template"

    [x] get rid of ./project-template/.tools/ ZSH-Scripts
        [x] startserver.zsh 
        [x] stopserver.zsh 
        [x] entr.zsh

    [ ] Create Script to remove Demo-Project files from `project-template`
    [ ] make it a `make` task


Loading Compressed Data:
================================================================================
    [x] Bitmap data
    [ ] mapData


Documentation:
================================================================================

## Usage:
    [x] Project-Setup 
    [x] Engine-Lifecycle
    [ ] Project - main.go
    [x] Scenes
    [ ] Scene-Transitions
    [ ] Drawing Stuff
    [x] Mouse-Input
    [ ] Keyboard-Input
    [ ] Requesting Assets from the Browser
    [ ] Audio-Output
    [ ] Storing Data in the Browser
    
## Reference:
    [x] Engine
    [x] Engine-State
    [ ] Keyboard
    [ ] Canvas
    [ ] TileMap
        [ ] Concept
        [ ] Map creation
        [ ] Map Editing
        [ ] Scrolling
    [x] TileSet 
    [x] Bitmap
    [ ] Bitmap Entities

Done:
================================================================================
    [x] Restructure how BlitBitmap receives its params @done (24.9.2023, 07:19:41) As they will become more and more complex
    [x] Implement Bitmap ClipingRect @done (24.9.2023, 16:36:16)
    [x] Printing Text from string to Screen 
    [x] TileSets and Maps
        [x] TileSets are Bitmaps plus some metadata to find clipping Rects @done (26.9.2023, 07:07:20)
        [x] TileMaps are lists of ClippingRects
        [x] Map Drawing
    [x] Define basic Print Text Functions

### Mouse Input:
    [x] Browser side
    [x] Go Side
        [x] Process Mouse Position @done (26.9.2023, 05:39:49)
        [x] Process Button clicks
    [x] Mouse is not available to Tick function due to being bound to the canvas (Bind mouse to Engine instead)

    [x] Documentation: Setup.md => "if you have entr as available shell command[s]"
    [x] Documentation: Readme.md => Link to Graphics does not work
