//Ajax autofil function
$(document).ready(function(){
    $("#team1").keyup(function(){
        var team1 = $("#team1").val();
        if (team1.length > 1) {
            //console.log("sending: ",team1);

            $.ajax({
                url: "/hth",
                method: "POST",
                data : {
                    team1: team1,
                },
                success: function(data){
                    $("#autofil-team1").html(data);
                    console.log("server response: ", data);
                }
            });
        }
        else {
            $('#autofil-team1').empty();
        }
    });
});

$(document).ready(function(){
    $("#team2").keyup(function(){
        var team2 = $("#team2").val();
        if (team2.length > 1) {
            //console.log("sending: ",team1);

            $.ajax({
                url: "/hth",
                method: "POST",
                data : {
                    team2: team2,
                },
                success: function(data){
                    $("#autofil-team2").html(data);
                    //console.log("server response: ", data);
                }
            });
        }
        else {
            $('#autofil-team2').empty();
        }
    });
});

function fillTeam1(val){
    document.getElementById('team1').value = val;
    $('#autofil-team1').empty();
}

function fillTeam2(val){
    document.getElementById('team2').value = val;
    $('#autofil-team2').empty();
}