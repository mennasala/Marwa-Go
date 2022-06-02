<html>
    <head>
        <title>colleges List</title>
    </head>
    <body>

        <h3>colleges List</h3>
        <a href="/colleges/add">Add</a>
        <table border="1">
            <tr>
                <th>COL_ID</th>
                <th>COL_ADD</th>
                <th>COL_TYPE</th>
            </tr>
            {{range .colleges}}
                <tr>
                    <td>{{.COL_ID}}</td>
                    <td>{{.COL_ADD}}</td>
                    <td>{{.COL_TYPE}}</td>
                    <td>Edit | Delete</td>
                    <td>
                        < a href="/colleges/edit?col_id={{.id}}">  Edit</a>
                        < a href="/colleges/delete?col_id={{.id}}"onclick="return confirm
                        ("are you sure?")>Delete</a>
                    </td>
                </tr>
            {{end}}
        </table>
    </body>
</html>