var teams = [
    {
        "name": "Toronto Raptors",
        "short": "TOR"
    },
    {
        "name": "Miami Heat",
        "short": "MIA"
    }
];

var TeamDropdown = React.createClass({
    getInitialState: function(){
        return {
            url:""
        }
    },
    getTeamData: function(){
        var that = this;
        $.ajax({
            url: this.state.url,
            type: "GET"
        }).done(function(data){
            that.displayGames(data);
        }).fail(function(err){
            console.log(err);
        });
    },
    displayGames: function(data){
        var gameRadioButtons = React.createClass({
            render: function(){
                var createRadioButton = function(obj, i){
                    var label = new Date(obj.Year, obj.Month, obj.Day).toString().split(" ").splice(0, 4).join(" ") + " with " + obj.Away + " at " + obj.Home;
                    return (
                        <div key={i}>
                            <input id={obj.GameId} type="radio" name="game" value={obj.GameId}/> {label}
                        </div>
                    )
                };

                return (
                    <form id="game_radio_buttons" action="">
                        {data.map(createRadioButton)} 
                    </form>  
                ) 
            }
        });
        var gameRadioButtonElement= React.createElement(
            gameRadioButtons, {}
        );
        ReactDOM.render(gameRadioButtonElement, document.getElementById("games"), function(){
            $("#game_radio_buttons").click(function(){
                var val = $("input[name='game']:checked", "#game_radio_buttons").val();
                console.log(val);
            });
        });
    },   
    selectFirstOption: function(e){
        var obj = {firstOption: e.target.value};
        var that = this;
        this.setState(obj, function(){
            this.setUrl();
        });
    },
    selectSecondOption: function(e){
        var obj = {secondOption: e.target.value};
        var that = this;
        this.setState(obj, function(){
            that.setUrl();
        });
    },
    setUrl: function(){
        var that = this;
        this.setState({url: "/api/" + this.state.firstOption + "/vs/" + this.state.secondOption }, function(){
            if (this.state.firstOption !== undefined && this.state.secondOption !== undefined){
                that.getTeamData.apply(that);
            }
        })
    },
    render: function(){
        var createSelect = function(obj, i){
            return <option key={i} id={obj.short} value={obj.short}>{obj.name}</option>;
        };

        return(
            <div>
                <div>
                    <select className="col-md-2" id="select-teams-1" onChange={this.selectFirstOption}>
                        <option value defaultValue>Select a team</option>
                        {teams.map(createSelect)}
                    </select>
                </div>
                <div>
                    <select className="col-md-2" id="select-teams-2" onChange={this.selectSecondOption}>
                        <option value defaultValue>Select a team</option>
                        {teams.map(createSelect)}
                    </select>
                </div>
                <div className="col-md-2" id="games"></div>
                <div className="clearfix"/>
            </div>
        );
    }
});

var MainContainer = React.createClass({
    render: function(){
        return (
            <div>
                <TeamDropdown/>
                <h1>Helloooo, world!</h1>   
            </div>
        );
    }
});


ReactDOM.render(
    <MainContainer/>,
    document.getElementById('example')
);
