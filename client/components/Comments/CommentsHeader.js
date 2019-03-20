import React from 'react'
import { Grid, Image } from 'semantic-ui-react'

const style = {
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
  headerLayout: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
  },
}

function CommentsHeader({ imgUrl, name, tag }) {
  return (
    <Grid style={style.header} columns={2}>
      <Grid.Column style={style.profilePicContainer} width={3}>
        <Image style={style.profilePic} src={imgUrl} />
      </Grid.Column>
      <Grid.Column style={style.headerLayout}>
        <Grid.Row>
          <h1>{name}</h1>
        </Grid.Row>
        <Grid.Row>
          <h4>{tag}</h4>
        </Grid.Row>
      </Grid.Column>
    </Grid>
  )
}

export default CommentsHeader
