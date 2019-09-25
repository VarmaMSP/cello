import Link from 'next/link'
import React, { Component } from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

interface Props {
  podcast: Podcast
}

export default class PodcastThumbnail extends Component<Props> {
  render() {
    const { podcast } = this.props

    return (
      <Link
        href={{ pathname: '/podcasts', query: { podcastId: podcast.id } }}
        as={`/podcasts/${podcast.id}`}
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
    )
  }
}
