$(document).ready(function () {
  $('#menu-images').addClass('active');
});

var TableRow = React.createClass({
  render: function() {
    var history = this.props.content;
    return (
        <tr key={this.props.index}>
          <td className="data-index">{history.Id.substring(0, 5)}</td>
          <td className="data-name no-wrap">{history.Tags}</td>
          <td className="data-name no-wrap">{history.Size && app.func.byteFormat(history.Size)}</td>
          <td className="data-name no-wrap">{app.func.relativeTime(new Date(history.Created * 1000))}</td>
          <td className="data-name no-wrap">{history.CreatedBy}</td>
        </tr>
    );
  }
});

var Table = React.createClass({
  getInitialState: function() {
    return {data: []};
  },
  componentDidMount: function() {
    var self = this, id = $('#image-id').val();
    app.func.ajax({type: 'GET', url: '/api/image/history/'+id, success: function (data) {
      self.setState({data: data});
    }});
  },
  render: function() {
    var rows = this.state.data.map(function(record, index) {
      return <TableRow key={record.name} index={index} content={record} />
    });
    return (
        <table className="table table-striped table-hover">
          <thead><tr><th>ID</th><th>Tags</th><th>Size</th><th>Created</th><th>CreatedBy</th></tr></thead>
          <tbody>{rows}</tbody>
        </table>
    );
  }
});

React.render(<Table />, document.getElementById('data'));