---
# Page settings
layout: default
keywords:
comments: false

# Hero section
title: Creating a multicomponent application with odo
description: Deploy a multicomponent application

# Micro navigation
micro_nav: true

# Page navigation
page_nav:
    prev:
        content: Creating a single-component application with odo
        url: '/docs/creating-a-single-component-application-with-odo'
    next:
        content: Creating an application with a database
        url: '/docs/creating-an-application-with-a-database'
---
`odo` allows you to create a multicomponent application, modify it, and
link its components in an easy and automated way.

This example describes how to deploy a multicomponent application - a
shooter game. The application consists of a front-end Node.js component
and a back-end Java component.

  - `odo` is installed.

  - You have a running OpenShift cluster. Developers can use [CodeReady
    Containers
    (CRC)](https://cloud.redhat.com/openshift/install/crc/installer-provisioned?intcmp=7013a000002CtetAAC)
    to deploy a local OpenShift cluster quickly.

  - Maven is installed.

# Creating a project

Create a project to keep your source code, tests, and libraries
organized in a separate single unit.

1.  Log in to an OpenShift cluster:
    
        $ odo login -u developer -p developer

2.  Create a project:
    
        $ odo project create myproject
         ✓  Project 'myproject' is ready for use
         ✓  New project created and now using project : myproject

# Deploying the back-end component

To create a Java component, import the Java builder image, download the
Java application and push the source code to your cluster with `odo`.

1.  Import `openjdk18` into the cluster:
    
        $ oc import-image openjdk18 \
        --from=registry.access.redhat.com/redhat-openjdk-18/openjdk18-openshift --confirm

2.  Tag the image as `builder` to make it accessible for odo:
    
        $ oc annotate istag/openjdk18:latest tags=builder

3.  Run `odo catalog list components` to see the created image:
    
        $ odo catalog list components
        Odo Supported OpenShift Components:
        NAME          PROJECT       TAGS
        nodejs        openshift     10,8,8-RHOAR,latest
        openjdk18     myproject     latest

4.  Create a directory for your components:
    
        $ mkdir my_components $$ cd my_components

5.  Download the example back-end application:
    
        $ git clone https://github.com/openshift-evangelists/Wild-West-Backend backend

6.  Change directory to the back-end source directory and check that you
    have the correct files in the directory:
    
        $ cd backend
        $ ls
        debug.sh  pom.xml  src

7.  Build the back-end source files with Maven to create a JAR file:
    
        $ mvn package
        ...
        [INFO] --------------------------------------
        [INFO] BUILD SUCCESS
        [INFO] --------------------------------------
        [INFO] Total time: 2.635 s
        [INFO] Finished at: 2019-09-30T16:11:11-04:00
        [INFO] Final Memory: 30M/91M
        [INFO] --------------------------------------

8.  Create a component configuration of Java component-type named
    `backend`:
    
        $ odo create openjdk18 backend --binary target/wildwest-1.0.jar
         ✓  Validating component [1ms]
         Please use `odo push` command to create the component with source deployed
    
    Now the configuration file `config.yaml` is in the local directory
    of the back-end component that contains information about the
    component for deployment.

9.  Check the configuration settings of the back-end component in the
    `config.yaml` file using:
    
        $ odo config view
        COMPONENT SETTINGS
        ------------------------------------------------
        PARAMETER         CURRENT_VALUE
        Type              openjdk18
        Application       app
        Project           myproject
        SourceType        binary
        Ref
        SourceLocation    target/wildwest-1.0.jar
        Ports             8080/TCP,8443/TCP,8778/TCP
        Name              backend
        MinMemory
        MaxMemory
        DebugPort
        Ignore
        MinCPU
        MaxCPU

10. Push the component to the OpenShift cluster.
    
        $ odo push
        Validation
         ✓  Checking component [6ms]
        
        Configuration changes
         ✓  Initializing component
         ✓  Creating component [124ms]
        
        Pushing to component backend of type binary
         ✓  Checking files for pushing [1ms]
         ✓  Waiting for component to start [48s]
         ✓  Syncing files to the component [811ms]
         ✓  Building component [3s]
    
    Using `odo push`, OpenShift creates a container to host the back-end
    component, deploys the container into a Pod running on the OpenShift
    cluster, and starts the `backend` component.

11. Validate:
    
      - The status of the action in odo:
        
            odo log -f
            2019-09-30 20:14:19.738  INFO 444 --- [           main] c.o.wildwest.WildWestApplication         : Starting WildWestApplication v1.0 onbackend-app-1-9tnhc with PID 444 (/deployments/wildwest-1.0.jar started by jboss in /deployments)
    
      - The status of the back-end component:
        
            $ odo list
            APP     NAME        TYPE          SOURCE                             STATE
            app     backend     openjdk18     file://target/wildwest-1.0.jar     Pushed

# Deploying the front-end component

To create and deploy a front-end component, download the Node.js
application and push the source code to your cluster with `odo`.

1.  Download the example front-end application:
    
        $ git clone https://github.com/openshift/nodejs-ex

2.  Change the current directory to the front-end directory:
    
        $ cd <directory-name>

3.  List the contents of the directory to see that the front end is a
    Node.js application.
    
        $ ls
        assets  bin  index.html  kwww-frontend.iml  package.json  package-lock.json  playfield.png  README.md  server.js
    
    <div class="note">
    
    The front-end component is written in an interpreted language
    (Node.js); it does not need to be built.
    
    </div>

4.  Create a component configuration of Node.js component-type named
    `frontend`:
    
        $ odo create nodejs frontend
         ✓  Validating component [5ms]
        Please use `odo push` command to create the component with source deployed

5.  Push the component to a running container.
    
        $ odo push
        Validation
         ✓  Checking component [8ms]
        
        Configuration changes
         ✓  Initializing component
         ✓  Creating component [83ms]
        
        Pushing to component frontend of type local
         ✓  Checking files for pushing [2ms]
         ✓  Waiting for component to start [45s]
         ✓  Syncing files to the component [3s]
         ✓  Building component [18s]
         ✓  Changes successfully pushed to component

# Linking both components

Components running on the cluster need to be connected in order to
interact. OpenShift provides linking mechanisms to publish communication
bindings from a program to its clients.

1.  List all the components that are running on the cluster:
    
        $ odo list
        APP     NAME         TYPE          SOURCE                             STATE
        app     backend      openjdk18     file://target/wildwest-1.0.jar     Pushed
        app     frontend     nodejs        file://./                          Pushed

2.  Link the current front-end component to the backend:
    
        $ odo link backend --port 8080
         ✓  Component backend has been successfully linked from the component frontend
        
        Following environment variables were added to frontend component:
        - COMPONENT_BACKEND_HOST
        - COMPONENT_BACKEND_PORT
    
    The configuration information of the back-end component is added to
    the front-end component and the front-end component restarts.

# Exposing components to the public

1.  Create an external URL for the application:
    
        $ cd frontend
        $ odo url create frontend --port 8080
         ✓  URL frontend created for component: frontend
        
        To create URL on the OpenShift  cluster, use `odo push`

2.  Apply the changes:
    
        $ odo push
        Validation
         ✓  Checking component [21ms]
        
        Configuration changes
         ✓  Retrieving component data [35ms]
         ✓  Applying configuration [29ms]
        
        Applying URL changes
         ✓  URL frontend: http://frontend-app-myproject.192.168.42.79.nip.io created
        
        Pushing to component frontend of type local
         ✓  Checking file changes for pushing [1ms]
         ✓  No file changes detected, skipping build. Use the '-f' flag to force the build.

3.  Open the URL in a browser to view the application.

<div class="note">

If an application requires permissions to the active Service Account to
access the OpenShift namespace and delete active pods, the following
error may occur when looking at `odo log` from the back-end component:

`Message: Forbidden!Configured service account doesn’t have access.
Service account may have been revoked`

To resolve this error, add permissions for the Service Account role:

    $ oc policy add-role-to-group view system:serviceaccounts -n <project>
    $ oc policy add-role-to-group edit system:serviceaccounts -n <project>

Do not do this on a production cluster.

</div>

# Modifying the running application

1.  Change the local directory to the front-end directory:
    
        $ cd ~/frontend

2.  Monitor the changes on the file system using:
    
        $ odo watch

3.  Edit the `index.html` file to change the displayed name for the
    game.
    
    <div class="note">
    
    A slight delay is possible before odo recognizes the change.
    
    </div>
    
    odo pushes the changes to the front-end component and prints its
    status to the terminal:
    
        File /root/frontend/index.html changed
        File  changed
        Pushing files...
         ✓  Waiting for component to start
         ✓  Copying files to component
         ✓  Building component

4.  Refresh the application page in the web browser. The new name is now
    displayed.

# Deleting an application

<div class="important">

Deleting an application will delete all components associated with the
application.

</div>

1.  List the applications in the current project:
    
        $ odo app list
            The project '<project_name>' has the following applications:
            NAME
            app

2.  List the components associated with the applications. These
    components will be deleted with the application:
    
        $ odo component list
            APP     NAME                      TYPE       SOURCE        STATE
            app     nodejs-nodejs-ex-elyf     nodejs     file://./     Pushed

3.  Delete the application:
    
        $ odo app delete <application_name>
            ? Are you sure you want to delete the application: <application_name> from project: <project_name>

4.  Confirm the deletion with `Y`. You can suppress the confirmation
    prompt using the `-f` flag.
