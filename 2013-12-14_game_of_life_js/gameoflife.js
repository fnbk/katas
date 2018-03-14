function isCellAlive(currentCellState, numberOfNeighbors){
    var aliveMap = {
        0: false,
        1: false,
        2: true,
        3: true,
        4: false,
        undefined: false
    };

    try {
        aliveMap[numberOfNeighbors].toString();
        return aliveMap[numberOfNeighbors];
    } catch(e) {
        return false;
    }
}
describe("A Cell", function() {
    it("should die if it has no alive neighbors", function() {
        var numberOfNeighbors = 0;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(false);
    });

    it("should die if it has one alive neighbor", function() {
        var numberOfNeighbors = 1;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(false);
    });

    it("should live if it has 2 alive neighbors", function() {
        var numberOfNeighbors = 2;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(true);
    });

    it("should live if it has 3 alive neighbors", function() {
        var numberOfNeighbors = 3;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(true);
    });

    it("should die if it has three neighbors", function() {
        var numberOfNeighbors = 4;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(false);
    });

    it("should die if it has EXACT three neighbors", function() {
        var numberOfNeighbors = 3;
        var currentCellState = false;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(true);
    });

    it("should die if it has more than three neighbors", function() {
        var numberOfNeighbors = 400;
        var currentCellState = true;
        var cellState = isCellAlive(currentCellState, numberOfNeighbors);
        expect(cellState).toBe(false);
    });
});