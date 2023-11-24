{{template "base.tpl"}}

<!-- Top navbar for xs viewports -->
<nav class="navbar navbar-inverse visible-xs">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">   
                <!-- navbar 3 horizontal lines -->
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">raven</a>
        </div>
        <div class="collapse navbar-collapse" id="myNavbar">
            <ul class="nav nav-pills nav-stacked">
                <li class="active"><a href="/">Home</a></li>
                <li><a href="#">Team Setup</a></li>
            </ul>
            <ul class="nav navbar-nav">
                {{range .teams}}
                <li><a href="#">{{.Name}}</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</nav>

<!-- Side navbar for non-xs viewports -->
<div class="container-fluid">
        <div class="col-sm-3 sidenav hidden-xs">
            <h2>raven</h2>
            <ul class="nav nav-pills nav-stacked">
                <li class="active"><a href="/">Home</a></li>
                <li><a href="/teams">Team Setup</a></li>
                <li><a href="/uploads">Uploads</a></li>
                <li><a href="/profile">Profile</a></li>
            </ul>
            <h4>Teams</h4>
            <ul class="nav nav-pills nav-stacked">
                {{range .teams}}
                <li><a href="/teams/{{.Id}}">{{.Name}}</a></li>
                {{end}}
            </ul>
            <br>
        </div>
<style>
    /* Set gray background color and 100% height */
    .sidenav {
      background-color: #f1f1f1;
      height: 100%;
    }
        
    /* On small screens, set height to 'auto' for the grid */
    @media screen and (max-width: 767px) {
      .row.content {height: auto;} 
    }

    html, body
    {
        height: 100%;
    }

    .container-fluid
    {
        height: 100%;
        /* overflow-y: hidden; */
    }   
  </style>

  <br>

  <div class="col-sm-9">
        {{.LayoutContent}}
  </div>
<br>

</div>
</body>
</html>
