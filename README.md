### Vault Demo

This is a simple example of how Hashicorp Vault could be used as the authentication backend to a web application.

The aspects of this demo include:

- [x] Web front end shows differences between authorized vs. unauthorized user access
- [x] Vault is used for user authentication
- [x] Vault userpass auth backend is used

##### Disclaimer

**<span style="color:red">NOTE</span>**: This is for demo purposes only.  For instance, the following aspects of this demo should be handled in a more secure way:

- Vault init payload should be encrypted using gpg keys or written to a secure data store
- Vault should use SSL certs that are not self-signed
- The webapp should use SSL so that the user and password provided on the web page are not transmitted in clear text
- The user credentials should be created within Vault using a more secure method then providing them in ```ansible/vault.yml```
- The password ```pass1``` is a weak password.  **Don't use weak passwords** (see [https://www.lastpass.com/password-generator](https://www.lastpass.com/password-generator) for help)

##### Other Notes

There is a [bug](https://github.com/TerryHowe/ansible-modules-hashivault/pull/191) within the python module ```ansible-modules-hashivault``` for when a wrapped token is created.  The ```ansible/requirements.txt``` file references an alternate repo since the [Pull Request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests) (PR) has not been merged as of the writing of this document.

#### Prerequisites

This demo uses Hashicorp Vagrant.  [Download](https://www.vagrantup.com/downloads.html) and [install](https://www.vagrantup.com/docs/installation/) Vagrant before continuing.

#### Walkthrough

1. Create a directory for this demo project and then go into this directory:

```sh
$ mkdir vaultapp
$ cd vaultapp
```

2. Download the Vagrant file to this project directory:

```sh
$ curl -O https://raw.githubusercontent.com/rascals77/vaultapp/master/Vagrantfile
```

3. Create the virtual machine:

```sh
$ vagrant up
```

4. Open a web browser and go to the following URL:

   [http://127.0.0.1:8080](http://127.0.0.1:8080)

5. Click on **Login** link in the top menu

6. Log in using **user1** for the **Username** and **pass1** for the **Password**

7. Notice the **Create Article** link appears in the top menu

8. Click on the **Create Article** link and create an article

9. Click on the **Logout** link in the top menu

10. Notice the **Create Article** link is no longer visible in the top menu

#### Clean up

To delete the virtual machine that was created, run the following command from the project directory:

```sh
$ vagrant destroy
```

