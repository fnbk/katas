var Cart = require('./checkout');


describe("total", function() {
    var cart;

    beforeEach(function() {
        cart = new Cart();
    });

    it("should be zero for new cart", function() {
        expect(cart.total()).toEqual(0);
    });

    it("should be non-zero for a cart with an item", function() {
        cart.addItem("A");
        expect(cart.total()).toBeGreaterThan(0);
    });

    it("should be 200 with only item B", function(){
        cart.addItem("B");
        expect(cart.total()).toEqual(200);
    });

    it("should ignore invalid item", function() {
        cart.addItem("A");
        var total = cart.total();
        expect(cart.total()).toEqual(total);
        cart.addItem("invalid");
        expect(cart.total()).toEqual(total);
    });

    it("should be 400 with two items B", function(){
        cart.addItems(["B", "B"]);
        expect(cart.total()).toEqual(400);
    });

    it("should be 300 after adding three As with package size three", function() {
        cart.addItems(["A", "A", "A"]);
        expect(cart.total()).toEqual(300);
    });

    xit("should be 800 after adding four Bs with package size four", function() {
        cart.addItems(["B", "B", "B", "B"]);
        expect(cart.total()).toEqual(800);
    });
});
