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
    selectFirstOption: function(e){
        var obj = {firstOption: e.target.value};
        this.setState(obj);
    },
    selectSecondOption: function(e){
        var obj = {secondOption: e.target.value};
        this.setState(obj);
    },
    viewState: function(){
        console.log(this.state);
    },
    render: function(){
        var createSelect = function(obj, i){
            return <option key={i} id={obj.short} value={obj.short}>{obj.name}</option>;
        };

        return(
            <div>
                <div className="row">
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
                    <div className="clearfix"/>
                    <button className="row col-md-2 col-md-offset-1" onClick={this.viewState}>Get Data</button>
                </div>
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
