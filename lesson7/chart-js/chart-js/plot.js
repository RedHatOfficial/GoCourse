$(document).ready(function() {
    $.getJSON("/values",
        function (data) {
            var graphElement = $('#graphElement')[0];
            var context = graphElement.getContext('2d');
            var myChart = new Chart(context, {
                type: 'bar',
                data: {
                    labels: data.x,
                    datasets: [{
                        label: '# of Votes',
                        data: data.y,
                        borderWidth: 1
                    }]
                },
                options: {
                    scales: {
                        yAxes: [{
                            ticks: {
                                beginAtZero: true
                            }
                        }]
                    },
                    maintainAspectRatio: false,
                }
            });
        });
});
