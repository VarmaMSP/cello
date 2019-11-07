import Grid from 'components/grid'
import Link from 'next/link'
import React, { Component } from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

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
          <Link
            href={{
              pathname: '/podcasts',
              query: { podcastId: podcast.id, activeTab: 'episodes' },
            }}
            as={`/podcasts/${podcast.id}/episodes`}
            key={podcast.id}
          >
            <div className="w-full cursor-pointer">
              <img
                className="w-full h-auto flex-none object-contain rounded-lg border"
                src={getImageUrl(podcast.id, 'md')}
              />
              <p className="text-xs tracking-wide leading-tight text-gray-800 mt-2 mb-1 line-clamp-2">
                {podcast.title}
              </p>
              <p className="text-xs tracking-tigher text-gray-600 truncate">
                {`by ${podcast.author}`}
              </p>
            </div>
          </Link>
        ))}
      </Grid>
    )
  }
}
