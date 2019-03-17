# Cloud Storage

#### Set Up
1. Make sure that you have set up an App Engine instance
2. Visit [this link](https://console.cloud.google.com/storage)
3. Select <project-name>.appspot.com and select `Edit bucket permissions` from the menu
4. In the `Add members` section, type in `allUsers` and select the `Storage` > `Storage Object Viewer`. This will make this whole bucket accessible to everyone in the world
5. Upload all of the documents in this directory

#### Why should we use this over hosting content ourselves?
For backend services, it is nice to have all of our servable content in a managed cloud provider so that no matter how many instances of our backend we have, we don't have to keep it in our backend instances.
If we had 100 instances and had to serve 1,000 resumes, that is a lot of duplicated data. Instead, we keep track of *where* our data is and communicate to the front end where to find it. When we host
content in Google Storage, we no longer have duplicated data across our instances. Uploading data to the server is also near impossible without a cloud provider. We cannot assume that we have permissions to write
to the host VM. We must write to Google Cloud Storage.
