const readline = require('readline'); // TODO use import statement; transpile with Babel

//
// asking CLI questions
//

function consoleQuestionPromise(question) {
    return new Promise(resolve => {
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });
        rl.question(`${question} `, (answer) => {
            resolve(answer);
            rl.close();
        });

        // TODO timeout
    });
}

// Iterate over functions, returning Promises and accumulate results in a collection
// source: http://bit.ly/2Bi5z7h
// Array.reduce() https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce
// Promise.then() https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then
const iterateSequentially = funcProms =>
    // callback(accumulator, currentValue, currentIndex, array), initialValue
    funcProms.reduce(
        // callback(accumulator, currentValue)
        (acc, cur) =>
            // Promise.resolve() returns a new Promise, since a value is passed, 'result' is the provided value
            // Promise.then() returns a Promise; so the next 'acc' is this new Promise
            acc.then(result =>
                // 'curr' is an element of the 'funcs' array; a function that returns a Promise
                // curr().then() returns a Promise
                cur().then( answer =>
                    // the Promise is resolved with the return value of the statement below ([1].concat(2) == [1,2])
                    result.concat(answer))),
        Promise.resolve([]) // initial value
        // To be consistent in the Promise cascade you can write the initial value as follows:
        // Promise.resolve(Promise.resolve([])) // initial value
    );
    // summary:
    // so essentially a collection of 'funcs' is reduced
    // each reduce step returns a new 'acc', which is a Promise, containing a Promise, containing the a collection


//
// Example: Promise.then() returns a Promise, containing a Promise, containing a value is resolved into Promise, containing a value
//

{
    // Promise.resolve(1)
    //     .then(a => Promise.resolve(2).then(b => a + b))
    //     .then(console.log.bind(console));
}


//
// helpers
//

var test = require('tape');
var tests = [];
const runTests = () => tests.forEach(t => t());

/**
 * Clones an object.
 *
 * @param {obj} the object to be cloned
 * @return returns a new object
 */
const clone = obj => Object.assign({}, obj);

tests.push(() => test('"clone" utility function test', function (t) {
    t.plan(1);
    let count = 0;
    times(2, () => count++);
    t.equal(count, 2);
}));

/**
 * Executes a callback n times.
 *
 * @param {number} the amount of times the callback needs to be called
 * @param {cb} the callback to be called
 */
const times = (number, cb) => Array.from(Array(number)).map( (v,i,a) => cb());

tests.push(() => test('"times" utility function test', function (t) {
    t.plan(1);
    let a = {};
    let b = clone(a);
    t.ok(a != b);
}));

const randomInt = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min;

const gameQuestion = (reiz) => `Reiz: ${reiz}, Antwort:`;

const calculateNBack = (stimuli, stimuliAnswers) => clone({stimuli, stimuliAnswers}); // TODO



//
// n-back
//

var startQuestions = [
    "Name des Probanden:",
    "n: Größe des Reiz-Abstandes nach dessen Übereinstimmung geprüft werden soll:",
    "Reizdauer: Wie lange soll jeder Reiz präsentiert werden? (Angabe in msec):",
    "Anzahl der Reize (10..100):",
];

const runNBack = () =>
    iterateSequentially(startQuestions.map(q => () => consoleQuestionPromise(q)))
        .then(game => clone({
            name: game[0],
            stimuliGap:  parseInt(game[1]),
            stimuliTime: parseInt(game[2]),
            stimuliNumber: parseInt(game[3])
        }))
        .then(game => clone({...game, stimuli: times(game.stimuliNumber, () => randomInt(0,9))}))
        .then(game => clone({...game, stimuliQuestions: game.stimuli.map(gameQuestion).map(question => () => consoleQuestionPromise(question))}))
        .then(game => iterateSequentially(game.stimuliQuestions).then( answers => clone({...game, stimuliAnswers: answers})))
        .then(game => calculateNBack(game.stimuli, game.stimuliAnswers))
        .then(console.log.bind(console));




//
// THIS IS THE ENTRY POINT
//

var argv = require('yargs')
    .usage('Usage: $0 [options]')
    .example('$0', 'run n-back')
    .example('$0 --test', 'run tests')
    .alias('t', 'test')
    .describe('test', 'run tests')
    .help('h')
    .alias('h', 'help')
    .argv;

argv.test ? runTests() : runNBack();

