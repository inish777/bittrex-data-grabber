<template>
    <div id="app">
        <div>
            USD/BTC: {{btc}} USD/LTC: {{ltc}} USD/ETH: {{eth}}
        </div>
        <div>
            <vue-highcharts :options="btcOptions" ref="btcCharts">
            </vue-highcharts>
            <vue-highcharts :options="ltcOptions" ref="ltcCharts">
            </vue-highcharts>
            <vue-highcharts :options="ethOptions" ref="ethCharts">
            </vue-highcharts>
        </div>
    </div>
</template>

<script>
import VueHighcharts from 'vue2-highcharts';

export default {
    components: {
        VueHighcharts
    },
    data() {
        return {
            btcOptions: {
                chart: {
                    type: 'line'
                },
                title: {
                    text: 'USD/BTC'
                },
                subtitle: {
                    text: 'Source: bittrex.com'
                },
                xAxis: {
                    type: 'datetime'
                },
                yAxis: {
                    labels: {
                        formatter: function() {
                            return this.value + '$';
                        }
                    }
                },
                tooltip: {
                    crosshairs: true,
                    shared: true
                },
                credits: {
                    enabled: false
                },
                series: [{
                    name: 'Last',
                    data: []
                }]
            },
            ltcOptions: {
                chart: {
                    type: 'line'
                },
                title: {
                    text: 'USD/LTC'
                },
                subtitle: {
                    text: 'Source: bittrex.com'
                },
                xAxis: {
                    type: 'datetime'
                },
                yAxis: {
                    labels: {
                        formatter: function() {
                            return this.value + '$';
                        }
                    }
                },
                tooltip: {
                    crosshairs: true,
                    shared: true
                },
                credits: {
                    enabled: false
                },
                series: [{
                    name: 'Last',
                    data: []
                }]
            },
            ethOptions: {
                chart: {
                    type: 'line'
                },
                title: {
                    text: 'USD/ETH'
                },
                subtitle: {
                    text: 'Source: bittrex.com'
                },
                xAxis: {
                    type: 'datetime',
                    legend: "Date and time"
                },
                yAxis: {
                    labels: {
                        formatter: function() {
                            return this.value + '$';
                        }
                    }
                },
                tooltip: {
                    crosshairs: true,
                    shared: true
                },
                credits: {
                    enabled: false
                },
                series: [{
                    name: 'Last',
                    data: []
                }]
            },
            btc: 0,
            ltc: 0,
            eth: 0
        };
    },
    methods: {
        getBtc() {
            fetch("/api/btc")
                .then((response) => response.json())
                .then((result) => {
                    this.btc = result.value;
                    let chart = this.$refs.btcCharts.getChart();
                    chart.series[0].addPoint([result.timestamp * 1000, result.value], true);
                });
        },
        getLtc() {
            fetch("/api/ltc")
                .then((response) => response.json())
                .then((result) => {
                    this.ltc = result.value;
                    let chart = this.$refs.ltcCharts.getChart();
                    chart.series[0].addPoint([result.timestamp * 1000, result.value], true);
                });
        },
        getEth() {
            fetch("/api/eth")
                .then((response) => response.json())
                .then((result) => {
                    this.eth = result.value;
                    let chart = this.$refs.ethCharts.getChart();
                    chart.series[0].addPoint([result.timestamp * 1000, result.value], true);
                });
        }
    },
    mounted() {
        setInterval(() => {
            this.getBtc();
            this.getLtc();
            this.getEth();
        }, 1000);
    }
}
</script>

<style  lang="scss" rel="stylesheet/scss" type="text/css" scoped>
$chartWidth: 300px;
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

html,
body {
    margin: 0;
    padding: 0;
    background: rgba(37, 37, 37, 0.84);
}

#app {}

h1 {
    font-family: Dosis, "Source Sans Pro", "Helvetica Neue", Arial, sans-serif;
    padding: 2em 0 1em;
    text-align: center;
    a {
        text-decoration: none;
        color: #ccc;
    }
}

h1,
h3 {
    color: #ccc;
    font-weight: 300;
    p {
        font-size: 14px;
        margin: 5px;
    }
}

.charts {
    width: $chartWidth;
    margin: 0 auto;
    padding-bottom: 2em;
    text-align: center;
    h3 {
        text-align: center;
        margin: 0;
    }
    button {
        padding: 5px 20px;
        background: #fff;
        -webkit-appearance: none;
        border: 1px solid #c5c5c5;
        border-radius: 5px;
        outline: none;
        &:hover {
            background: lightgray;
            color: #333;
        }
    }
}

.highcharts {
    display: inline-block;
    margin: 2em auto;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    box-shadow: 0 0 45px rgba(0, 0, 0, 0.2);
    padding: 1.5em 0em;
    background: #fff;
    width: 30%;
}
</style>
