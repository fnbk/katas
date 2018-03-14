var Cart = function Cart() {
    //var total = 0;

    var itemCount = {
        "A": 0,
        "B": 0
    };

    var priceMatrix = {
        "A" : 150,
        "B" : 200
    };

    var discountPriceMatrix = {
        "A" : 100,
        "B" : 150
    };

    return {
        total: function () {
            var sum = 0;
            for (var key in itemCount) {
                var numberOfPackages = Math.floor(itemCount[key] / 3);
                var normalItems = itemCount[key] % 3;

                var packagePrice = numberOfPackages * discountPriceMatrix[key] * 3;
                var normalPrice = normalItems * priceMatrix[key];

                sum += packagePrice + normalPrice;
            }
            return sum;
        },

        addItem: function (item) {
            if (!priceMatrix[item]) return;
            itemCount[item] += 1;
        },

        addItems: function(items){
            for (var i = 0; i < items.length; i++){
                this.addItem(items[i]);
            }
        }
    }
};

module.exports = Cart;
