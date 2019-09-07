import React, { Component } from 'react'
import { Podcast } from 'types/app'
// import Link from 'next/link'
import ResponsiveGrid from '../responsive_grid'
import PodcastThumbnail from '../../components/podcast_thumbnail'

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
      <ResponsiveGrid rows={-1}>
        {podcasts.map((podcast) => (
          <PodcastThumbnail podcast={podcast} />
        ))}
      </ResponsiveGrid>
    )
  }
}
