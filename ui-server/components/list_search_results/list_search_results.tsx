import Grid from 'components/grid'
import PodcastThumbnail from 'components/podcast_thumbnail'
import React, { Component } from 'react'
import { Podcast } from 'types/app'

export interface StateToProps {
  podcasts: Podcast[]
}

export interface OwnProps {
  query: string
}

interface Props extends StateToProps, OwnProps {}

export default class extends Component<Props> {
  render() {
    const { podcasts } = this.props
    return (
      <Grid totalRowSpacing={{ LG: 12, MD: 8, SM: 8 }} className="mb-3 pb-3">
        {podcasts.map((podcast) => (
          <PodcastThumbnail key={podcast.id} podcast={podcast} />
        ))}
      </Grid>
    )
  }
}
