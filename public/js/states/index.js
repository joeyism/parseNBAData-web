var TeamDropdown = React.createClass({
    selectOption: function(e){
        var obj = {value: e.target.value};
        this.setState(obj);
        console.log(obj);
    },
    render: function(){
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
        var createSelect = function(obj, i){
            return <option key={i} id={obj.short} value={obj.short}>{obj.name}</option>;
        };
        
        return(
            <div>
            <select id="select-teams" onChange={this.selectOption}>
            <option disabled defaultValue value>Select a team</option>
            {teams.map(createSelect)}
            </select>
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
