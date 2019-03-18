import React from 'react'
import { Grid, Form, TextArea, Image } from 'semantic-ui-react'

const style = {
  form: {
    height: '100%',
    background: 'white',
    margin: '25px 0px',
    borderRadius: '50px 0px 0px 50px',
    borderTop: '50px solid #595a4a',
  },
  profilePicContainer: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  profilePic: {
    height: '100px',
    width: '100px',
    borderRadius: '50px',
    alignSelf: 'center',
    border: '4px solid #595a4a',
  },
  header: {
    margin: '4px',
  },
}

function TextEditor() {
  return (
    <Form style={style.form}>
      <Grid style={style.header} columns={2}>
        <Grid.Column style={style.profilePicContainer} width={3}>
          <Image style={style.profilePic} src='https://react.semantic-ui.com/images/wireframe/image.png' />
        </Grid.Column>
        <Grid.Column>
          <Grid.Row>
            <h1>Googler</h1>
          </Grid.Row>
          <Grid.Row>
            <h4>What do you think about this resume?</h4>
          </Grid.Row>
        </Grid.Column>
      </Grid>
      <TextArea placeholder='Please be nice to your friends' />
    </Form>
  )
}

export default TextEditor
