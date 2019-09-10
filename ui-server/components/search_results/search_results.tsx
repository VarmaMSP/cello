import PodcastThumbnail from 'components/podcast_thumbnail'
import ResponsiveGrid from 'components/responsive_grid'
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
      <ResponsiveGrid>
        {podcasts.map((podcast) => (
          <PodcastThumbnail key={podcast.id} podcast={podcast} />
        ))}
      </ResponsiveGrid>
    )
  }
}
