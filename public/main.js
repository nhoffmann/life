let params = (new URL(document.location)).searchParams;
let pattern = params.get('pattern');
let rule = params.get('rule');

if (!pattern) {
  pattern = 'RPentomino'
}

if (!rule) {
  rule = 'B3/S23';
}

window.pattern = pattern;
window.rule = rule;

function reload(event) {
  location.href = '?pattern='+event.target.value
}

document.getElementById('pattern-selector').addEventListener('change', reload)
document.getElementById('pattern-label').innerHTML = window.pattern

function readPatternFile(event) {
  const reader = new FileReader()
  reader.onload = function() {
    loadPatternFile(reader.result)
  }
  reader.readAsText(event.target.files[0])
}

document.getElementById('file-input').addEventListener('change', readPatternFile)

document.getElementById('center').addEventListener('click', function() {
  center()
})

document.getElementById('cellSideLength-range').addEventListener('input', function() {
  cellSideLength(event.target.value)
})

let startX = -1
let startY = -1

function throttled(delay, fn) {
  let lastCall = 0;
  return function (...args) {
    const now = (new Date).getTime();
    if (now - lastCall < delay) {
      return;
    }
    lastCall = now;
    return fn(...args);
  }
}

function throttledMoveCanvas(event) {
  throttled(20, moveCanvas)
}

function moveCanvas(event) {
  const now = (new Date).getTime();
  if (startX === -1 && startY === -1) {
    startX = event.clientX
    startY = event.clientY
  }

  setOffset((event.clientX - startX), (event.clientY - startY))

  startX = event.clientX
  startY = event.clientY
}

const init = function() {
  setInterval(function() {
    document.getElementById('cell-count').innerHTML = cellCount() || 0
    document.getElementById('generation-count').innerHTML = generationCount() || 0
  })

  const $canvas = document.getElementsByTagName('canvas')[0]
  $canvas.addEventListener('mousedown', function() {
    $canvas.addEventListener('mousemove', moveCanvas)
  })
  $canvas.addEventListener('mouseup', function() {
    $canvas.removeEventListener('mousemove', moveCanvas)
    startX = -1
    startY = -1
  })
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch('main.wasm'),go.importObject).then( res => {
  go.run(res.instance)

  init()
})
