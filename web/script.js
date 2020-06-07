class State {
  constructor(stateName, color) {
    this.name = stateName;
    this.color = color;
  }
}

var gradi = 19;
var max = 34;
var min = 2;

var stateIndex = -1;
var states = [new State("HEATING", "#fe4800"), new State("cooling", "#007af1"), new State("OFF", "white")];

async function loadTemperature(){
  let host = document.location.hostname
  let port = document.location.port
  let responce = await fetch("http://"+host+":"+port+"/temperature/current");
  let data = await responce.text();
  return data;
}

function changeStatus() {
  ++stateIndex;
  stateIndex = (stateIndex%states.length);
  document.getElementById("state").innerHTML = states[stateIndex].name;
  document.getElementById("state").style.color = states[stateIndex].color;
}

function decrees() {
  if (gradi > min) {
    gradi--;
    updateGr();
    if (gradi >= 18) {
      var fill1 = document.getElementsByClassName("fill1");
      for (i = 0 ; i < fill1.length; i++) {
        fill1[i].style.transform = "rotate(" + (gradi - 18) * 10 + "deg)";
        fill1[i].style.transitionDelay = "0s";
      }

    } else if (gradi == 17) {
      var fill2 = document.getElementsByClassName("fill2");
      for (i = 0 ; i < fill2.length; i++) {
        fill2[i].style.transform = "rotate(" + gradi * 10 + "deg)";
        fill2[i].style.transitionDelay = "0.5s";
      }

    } else {
      var fill2 = document.getElementsByClassName("fill2");
      for (i = 0 ; i < fill2.length; i++) {
        fill2[i].style.transform = "rotate(" + gradi * 10 + "deg)";
        fill2[i].style.transitionDelay = "0s";
      }

    }
  }
};

function increase() {
  if (gradi < max) {
    gradi++;
    updateGr();
    if (gradi > 19) {
      var fill1 = document.getElementsByClassName("fill1");
      for (i = 0 ; i < fill1.length; i++) {
        fill1[i].style.transform = "rotate(" + (gradi - 18) * 10 + "deg)";
        fill1[i].style.transitionDelay = "0s";
      }

    } else if (gradi == 19) {
      var fill1 = document.getElementsByClassName("fill1");
      for (i = 0 ; i < fill1.length; i++) {
        fill1[i].style.transform = "rotate(" + (gradi - 18) * 10 + "deg)";
        fill1[i].style.transitionDelay = "1s";
      }

    } else {
      var fill2 = document.getElementsByClassName("fill2");
      for (i = 0 ; i < fill2.length; i++) {
        fill2[i].style.transform = "rotate(" + gradi * 10 + "deg)";
        fill2[i].style.transitionDelay = "0s";
      }
    }
  }
};

function updateCurrentGr(temperature){
  var heat = document.getElementsByClassName("heat");
  for (i=0; i<heat.length; i++) {
    heat[i].innerHTML = ("" + parseFloat(temperature).toFixed(2));
  }
}

function updateGr() {
  var ext = document.getElementsByClassName("ext");
  for (i=0; i<ext.length; i++) {
    ext[i].innerHTML = ("" + gradi);
  }

  var numberElements = document.getElementsByClassName("number");
  for (i=0; i<numberElements.length; i++) {
    numberElements[i].style.transform = "translate(-50%, -50%) rotate(" + (-180 + gradi * 10) + "deg)";
  }

  var shadowElements = document.getElementsByClassName("shadow");
  for(i=0; i<shadowElements.length; i++) {
    shadowElements[i].style.transform = "translate(-50%, -50%) rotate(" + (-180 + gradi * 10) + "deg)";
    shadowElements[i].style.animation = "none";
  }

  var fillElements = document.getElementsByClassName("fill");
  for (i=0 ; i < fillElements.length; i++) {
    fillElements[i].style.animation = "none";
  }
};