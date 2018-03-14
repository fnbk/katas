class Cell
  def initialize(state)
    @state = state
  end

  def neighbors(amount)
    @state = :dead if amount < 2 || amount > 3
    @state = :live if amount == 3
  end

  def iterate(&health_check)
    health_check.call(@state)
  end
end

class Board
  def iterate!

  end

  def living_cells(&health_check)
    health_check.call([])
  end
end