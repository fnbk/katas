require_relative 'game'

describe "a cell" do
  it "dies when it has fewer than 2 active neighbors" do
    cell = Cell.new(:live)
    cell.neighbors(1)
    cell.iterate {|state| state.should == :dead }
  end

  it "stays alive when it has 2 active neighbors" do
    cell = Cell.new(:live)
    cell.neighbors(2)
    cell.iterate {|state| state.should == :live }
  end

  it "dies when it has more than 3 active neighbors" do
    cell = Cell.new(:live)
    cell.neighbors(4)
    cell.iterate {|state| state.should == :dead }
  end

  it "gets born when it has 3 active neighbors" do
    cell = Cell.new(:dead)
    cell.neighbors(3)
    state = :dead
    cell.iterate {|s| state = s }
    state.should == :live
  end


end

describe "a board" do
  it "that is empty stays empty" do
    board = Board.new
    board.iterate!
    board.living_cells do |cells|
      cells.should be_empty
    end
  end

  xit "that has an active cell on position 0x0 becomes empty" do
    board = Board.new([[-1, 1]])
    board.iterate!
    board.living_cells do |cells|
      cells.should be_empty
    end
  end
end