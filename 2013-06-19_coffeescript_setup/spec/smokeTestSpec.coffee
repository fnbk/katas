Smoke = require '../lib/smoke'

describe "smoke test", ->
  it "should be true", ->
    expect(true).toBe true

  it "should exist", ->
    smoke = new Smoke()
    expect(smoke.test()).toBe true
