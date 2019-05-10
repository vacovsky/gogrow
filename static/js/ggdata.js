// "use strict";

// var dumpData = {};


// Chart.defaults.global.defaultColor = rgba(0, 0, 0, 0.4);

function loadDump() {
    var Http = new XMLHttpRequest();
    var url = "/dump";

    Http.open("GET", url);
    Http.responseType = "json";
    Http.setRequestHeader("Content-Type", "application/json");

    Http.send();

    Http.onreadystatechange = (e) => {
        var dumpData = Http.response;
        if (dumpData != null) {
            populateTemplate(dumpData);
        }
    };
}

function CtoF(n) {
    return (n * (9 / 5) + 32);
}

function windDegConverter(deg) {
    var cdIndex = parseInt((deg / 22.5) + .5, 10);
    var arr = ["N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"];
    return arr[(cdIndex % 16)];
}

function populateTemplate(data) {
    weld(document.querySelector("#gg-weather-description"), data.Weather.Description + " / " + data.Weather.Clouds + "% cloudiness");
    weld(document.querySelector("#gg-weather-label"), "Conditions as of " + moment(data.Weather.TimeStamp).fromNow());
    weld(document.querySelector("#gg-temp-humidity-value"), data.Weather.Temperature + "°F / " + data.Weather.Humidity + "%");
    weld(document.querySelector("#gg-tent-temp-hum"), CtoF(data.Device.TentTemperature) + "°F / " + data.Device.TentHumidity + "%");
    weld(document.querySelector("#gg-ambient-temp-hum"), CtoF(data.Device.AmbientTemperature) + "°F / " + data.Device.AmbientHumidity + "%");
    weld(document.querySelector("#gg-wind-speed-direction"), data.Weather.WindSpeed + " mph / " + windDegConverter(data.Weather.WindDeg));
    weld(document.querySelector("#gg-air-pressure"), data.Weather.AirPressure + " mbar");

    weld(document.querySelector("#gg-grow-cycle-plant"), data.Cycle.FriendlyName);
    weld(document.querySelector("#gg-grow-cycle-sewn"), moment(data.Cycle.TimeOfPlanting).fromNow());
    weld(document.querySelector("#gg-grow-cycle-germinated"), moment(data.Cycle.TimeGerminated).fromNow());
    weld(document.querySelector("#gg-grow-cycle-harvested"), moment(data.Cycle.TimeHarvested).fromNow());
    weld(document.querySelector("#gg-grow-cycle-flowered"), moment(data.Cycle.TimeOfFlowering).fromNow());
    weld(document.querySelector("#gg-grow-cycle-days"), moment(data.Cycle.TimeOfPlanting).fromNow());

    weld(document.querySelector("#gg-watered-time-volume"), data.Water.AmountUsedLiters + "L / " + moment(data.Water.Time).fromNow());
    weld(document.querySelector("#gg-fertilizer-time"), moment(data.Fertilization.Time).fromNow());
    weld(document.querySelector("#gg-fertilizer-type-amount"), data.Fertilization.AmountUsedLiters + "L");
    weld(document.querySelector("#gg-fertilizer-nkp"), data.Fertilization.FertilizerN + " / " + data.Fertilization.FertilizerK + " / " + data.Fertilization.FertilizerP);

    weld(document.querySelector("#gg-fertilizer-type-amount"), data.Fertilization.AmountUsedLiters + "L");
    weld(document.querySelector("#gg-fertilizer-type-amount"), data.Fertilization.AmountUsedLiters + "L");

    weld(document.querySelector("#gg-light-type"), data.Light.Type);
    weld(document.querySelector("#gg-light-wattage"), data.Light.Wattage + " Watts");

    weld(document.querySelector("#gg-fan-type"), data.Fan.Type);
    weld(document.querySelector("#gg-fan-cfm"), data.Fan.CFM + " CFM");

}

$(document).ready(function () {
    console.log("ready!");
});

function submitWaterForm() {
    var waterFormData = {
        AmountUsedLiters: parseInt($('#WaterAmountUsedLiters').val(), 10) || 1,
        Time: $('#WaterTime').val() + ":00Z" || moment().unix()
    };
    $.ajax({
        type: "POST",
        url: "/water",
        data: JSON.stringify(waterFormData),
        success: function () {
            loadDump();
            $.featherlight.close();
        },
        contentType: "application/json"
    });
}


function submitFanForm() {
    var waterFormData = {
        CFM: parseInt($('#FanCFM').val(), 10),
        Type: $('#FanType').val()
    };
    $.ajax({
        type: "POST",
        url: "/fan",
        data: JSON.stringify(waterFormData),
        success: function () {
            loadDump();
            $.featherlight.close();
        },
        contentType: "application/json"
    });
}


function submitLightForm() {
    var formData = {
        Wattage: parseInt($('#LightWattage').val(), 10),
        Type: $('#LightType').val() || "unknown",
    };
    $.ajax({
        type: "POST",
        url: "/light",
        data: JSON.stringify(formData),
        success: function () {
            loadDump();
            $.featherlight.close();
        },
        contentType: "application/json"
    });
}

function submitFertForm() {
    var formData = {
        AmountUsedLiters: parseInt($('FertAmountUsedLiters').val(), 10),
        Time: $('#FertTime').val() + ":00Z" || moment().unix(),
        FertType: parseInt($('FertType').val(), 10),
        FertilizerN: parseInt($('FertilizerN').val(), 10),
        FertilizerK: parseInt($('FertilizerK').val(), 10),
        FertilizerP: parseInt($('FertilizerP').val(), 10)
    };

    $.ajax({
        type: "POST",
        url: "/fertilization",
        data: JSON.stringify(formData),
        success: function () {
            loadDump();
            $.featherlight.close();
        },
        contentType: "application/json"
    });
}

function submitCycleForm() {
    var formData = {
        FriendlyName: $('#FriendlyName').val(),
        TimeOfPlanting: $('#TimeOfPlanting').val() + ":00Z",
        TimeGerminated: $('#TimeGerminated').val() + ":00Z",
        TimeOfFlowering: $('#TimeOfFlowering').val() + ":00Z",
        TimeHarvested: $('#TimeHarvested').val() + ":00Z",
        PlantedAsSeed: $('#PlantedAsSeed').val() == "on"
    };
    $.ajax({
        type: "POST",
        url: "/cycle",
        data: JSON.stringify(formData),
        success: function () {
            loadDump();
            $.featherlight.close();
        },
        contentType: "application/json"
    });
}

var CHART_OPTIONS = {

    scales: {
        xAxes: [{
            type: 'time',
            time: {
                unit: 'month'
            }
        }]
    }

};


function loadTempChart() {

    var Http = new XMLHttpRequest();
    var url = "/chart/temp";

    Http.open("GET", url);
    Http.responseType = "json";
    Http.setRequestHeader("Content-Type", "application/json");

    Http.send();

    Http.onreadystatechange = (e) => {
        var chartData = Http.response;

        if (chartData != undefined) {

            console.log(chartData);
            var config = {
                type: 'line',
                data: {
                    labels: chartData.Labels,
                    datasets: [{
                            label: "Tent",
                            borderColor: "lime",
                            data: chartData.Data[0],
                            fill: false,
                        },
                        {
                            label: "Ambient",
                            fill: false,
                            borderColor: window.chartColors.orange,
                            data: chartData.Data[1],
                        },
                        {
                            label: "Outside",
                            fill: false,
                            borderColor: window.chartColors.blue,
                            data: chartData.Data[2],
                        },
                        {
                            label: "CPU",
                            fill: false,
                            borderColor: window.chartColors.red,
                            data: chartData.Data[3],
                        }
                    ]
                },
                options: {
                    legend: {
                        labels: {
                            fontColor: "#bbb",
                        }
                    },
                    elements: {
                        point: {
                            radius: 0
                        }
                    },
                    aspectRatio: 2,
                    responsive: true,
                    title: {
                        display: true,
                        text: 'Temperature'
                    },
                    tooltips: {
                        mode: 'index',
                        intersect: false,
                    },
                    hover: {
                        mode: 'nearest',
                        intersect: true
                    },
                    scales: {
                        xAxes: [{
                            type: 'time',
                            display: true,
                            ticks: {
                                fontColor: "#bbb",
                            },
                            scaleLabel: {
                                display: true,
                            }
                        }],
                        yAxes: [{
                            display: true,
                            scaleLabel: {
                                display: true,
                                labelString: '°F'
                            },
                            ticks: {
                                beginAtZero: true,
                                min: 0,
                                max: 120,
                                fontColor: "#bbb",
                                stepSize: 20,
                            }
                        }]
                    }
                }
            };


            var ctx = document.getElementById('temp-chart').getContext('2d');

            new Chart(ctx, config);
        }
    };
}



function loadHumChart() {

    var Http = new XMLHttpRequest();
    var url = "/chart/hum";

    Http.open("GET", url);
    Http.responseType = "json";
    Http.setRequestHeader("Content-Type", "application/json");
    Http.send();

    Http.onreadystatechange = (e) => {
        var chartData = Http.response;


        if (chartData != undefined) {
            console.log(chartData);
            var config = {
                type: 'line',
                data: {
                    labels: chartData.Labels,
                    datasets: [{
                            label: "Tent",
                            fill: false,
                            borderColor: "lime",
                            data: chartData.Data[0]
                        },
                        {
                            label: "Ambient",
                            fill: false,
                            borderColor: window.chartColors.orange,
                            data: chartData.Data[1]
                        },
                        {
                            label: "Outside",
                            fill: false,
                            borderColor: window.chartColors.blue,
                            data: chartData.Data[2]
                        },
                        {
                            label: "Cloud Coverage",
                            fill: false,
                            borderColor: "#aaa",
                            data: chartData.Data[3]
                        }
                    ]
                },
                options: {
                    legend: {
                        labels: {
                            fontColor: "#bbb",
                            // fontSize: 18
                        }
                    },
                    elements: {
                        point: {
                            radius: 0
                        }
                    },
                    aspectRatio: 2,
                    responsive: true,
                    title: {
                        display: true,
                        text: 'Humidity & Clouds'
                    },
                    tooltips: {
                        mode: 'index',
                        intersect: false,
                    },
                    hover: {
                        mode: 'nearest',
                        intersect: true
                    },
                    scales: {
                        xAxes: [{
                            type: 'time',
                            display: true,
                            ticks: {
                                fontColor: "#bbb",
                            },
                            scaleLabel: {
                                display: true,
                                // labelString: 'Month'
                            },
                        }],
                        yAxes: [{
                            display: true,
                            scaleLabel: {
                                display: true,
                                labelString: '%'
                            },
                            ticks: {
                                beginAtZero: true,
                                min: 0,
                                max: 100,
                                stepSize: 20,
                                fontColor: "#bbb",
                            }
                        }]
                    }
                }
            };

            var ctx = document.getElementById('hum-chart').getContext('2d');
            new Chart(ctx, config);
        }
    };
}


loadDump();
loadTempChart();
loadHumChart();