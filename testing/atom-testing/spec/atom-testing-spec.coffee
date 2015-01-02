{WorkspaceView} = require 'atom'
AsciiArt = require '../lib/ascii-art'

# Use the command `window:run-package-specs` (cmd-alt-ctrl-p) to run specs.
#
# To run a specific `it` or `describe` block add an `f` to the front (e.g. `fit`
# or `fdescribe`). Remove the `f` to unfocus the block.

describe "atom-testing", ->
  promise = null
  beforeEach ->
    atom.workspaceView = new WorkspaceView()
    atom.workspace = atom.workspaceView.model
    promise = atom.packages.activatePackage('atom-testing')
    waitsForPromise ->
      atom.workspace.open()

  it "tests", ->
    atom.workspaceView.trigger 'atom-testing:test'
    waitsForPromise ->
      promise
