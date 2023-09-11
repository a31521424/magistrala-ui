
## Mainflux-ui

The Mainflux-ui functions as a Graphical User Interface (GUI) designed to interact with Mainflux services, encompassing both the creation and management aspects of users, things, channels, and groups. It streamlines tasks such as user and thing creation, channel establishment, policy configuration, and HTTP message transmission across various channels.

Mainflux-ui can be obtained as an independent subset of Mainflux; however, it requires integration with the mainflux core services to function properly.

## Prerequisites

To run Mainflux-ui, you need the following components:

- [Mainflux](https://github.com/mainflux/mainflux) (latest version)
- [Go](https://golang.org/doc/install) (version 1.19.2)
- [Docker-compose](https://docs.docker.com/compose/install/) (latest version)
- [make](https://www.gnu.org/software/make/manual/make.html)

## Installation

After installing the prerequisites, execute the following commands from the project's root directory:

```bash
make
```

```bash
make run
```

These commands will launch Mainflux-ui. To use the Mainflux-ui, ensure that the Mainflux core services are operational. You can achieve this by installing [Mainflux](https://github.com/mainflux/mainflux) and its prerequisites.

To build the docker images for the ui service, run the following commands which will build the docker images in different configurations.

This command will build the docker images for the ui service in default configuration.
```bash
make docker
```

This will build the development docker images for ui.
```bash
make docker_dev
```

You can also run ui via docker using the following command. 

```bash
make run_docker
```

This brings up the docker images and runs ui in the configuration specified in the .env file.

## Usage

Once both Mainflux core and Mainflux-ui are running, you can access the Mainflux-ui interface locally by entering the address: [http://localhost:9090](http://localhost:9090).

On the login page, use the provided credentials to log in to the interface:

```
Email: admin@example.com
Password: 12345678
```

Upon logging in, you will be directed to the Dashboard, which provides an overview of the Mainflux user interface. The sidebar elements, such as Users/Groups, allow you to navigate to specific pages for performing actions related to Users, Groups, Users Policies, Things, Channels, Things Policies, and viewing Deleted Clients.

### Users

You can create individual users or upload a CSV file to add multiple users. When creating a user, input the User Identity, User Secret, Tags (as a string slice), and Metadata (in JSON format). The User Identity should be unique and can be an email. The User Secret serves as a password for user login. Metadata provides additional user information.

When using a CSV file to create multiple users, the file should contain user names in one column and their respective credentials in subsequent columns. In the CSV file, the credentials should also follow in the order of email and then password.

### Groups

Similar to users, you can create single or multiple groups by uploading a CSV file. For group creation, provide the Group Name (required), Description, Parent ID, and Metadata (in JSON format).

When using a CSV file to create multiple groups, the file should contain group names in one column and corresponding Parent IDs in subsequent columns.

For more details, refer to the official [Documentation](http://docs.mainflux.io/cli/#things-management).

### Users Policies

To create a user policy, select the subject and object, and choose the relevant actions (multiple selections are allowed). User policies are utilized to manage permissions for users and groups entities. These Policies determine access rights for these entities. For instance, they define which users can access specific groups. Learn more about policies from the official [Documentation](https://docs.mainflux.io/authorization/#summary-of-defined-policies).

### Things

You can create individual things or upload a CSV file for multiple things. When adding a thing, provide the Thing Name (required), Thing ID, Identity, Secret, Tags (as a string slice), and Metadata (in JSON format). The Thing Secret should be unique and is used to identify the thing. Metadata offers additional information about the thing.

For multiple things, use a CSV file with thing names in one column. Refer to the official [Documentation](http://docs.mainflux.io/cli/#things-management) for CSV file details.

### Channels

Similarly, you can add individual or multiple channels using a CSV file. For channel creation, enter the Channel Name (required), select the Parent ID, provide a Description, and include Metadata (in JSON format).

### Things Policies

Creating a thing policy involves selecting the subject and object and specifying the actions (multiple selections are allowed). Things policies control permissions for things and channels entities. They define access permissions, such as a thing's access to a channel. For detailed policy information, consult the official [Documentation](https://docs.mainflux.io/authorization/#summary-of-defined-policies).

### Bootstrap

To use bootstrap, ensure that the [bootstrap](http://docs.mainflux.io/bootstrap/) addon is active as part of the Mainflux core services.

To configure bootstrap, provide the Name, Thing ID, External ID, External Key, Channel (as a string slice), Content (in JSON format), Client Cert, Client Key, and CA Cert.

