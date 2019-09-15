import Link from 'next/link'
import React, { Component } from 'react'
import Shiitake from 'shiitake'
import { Podcast } from 'types/app'

interface Props {
  podcast: Podcast
}

export default class PodcastThumbnail extends Component<Props> {
  render() {
    const { podcast } = this.props

    return (
      <Link
        href={{ pathname: '/podcasts', query: { id: podcast.id } }}
        as={`/podcasts/${podcast.id}`}
      >
        <div className="w-full cursor-pointer">
          <img
            className="w-full h-auto flex-none object-contain rounded-lg border"
            src={`${process.env.IMAGE_BASE_URL}/${podcast.id}p-500x500.jpg`}
          />
          <Shiitake
            lines={2}
            throttleRate={200}
            renderFullOnServer
            className="text-sm tracking-tight text-gray-800 mt-3 mb-1"
          >
            {podcast.title}
          </Shiitake>
          <Shiitake
            lines={1}
            throttleRate={200}
            renderFullOnServer
            className="text-sm tracking-tigher text-gray-600"
          >
            {`by ${podcast.author}`}
          </Shiitake>
        </div>
      </Link>
    )
  }
}
