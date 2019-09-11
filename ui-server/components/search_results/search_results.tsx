import Grid from 'components/grid'
import PodcastThumbnail from 'components/podcast_thumbnail'
import React, { Component } from 'react'
import { Podcast } from 'types/app'

export interface StateToProps {
  podcasts: Podcast[]
}

export interface OwnProps {
  searchQuery: string
}

interface Props extends StateToProps, OwnProps {}

export default class SearchResults extends Component<Props> {
  render() {
    const { podcasts } = this.props
    return (
      <Grid rowSpacing={12}>
        {podcasts.map((podcast) => (
          <PodcastThumbnail key={podcast.id} podcast={podcast} />
        ))}
      </Grid>
    )
  }
}
