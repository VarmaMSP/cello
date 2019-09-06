import React, { Component } from 'react'
import { Podcast } from 'types/app'
import Link from 'next/link'

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
      <div className="flex flex-wrap">
        {podcasts.map((podcast) => (
          <div key={podcast.id}>
            <Link
              href={{ pathname: '/podcasts', query: { id: podcast.id } }}
              as={`/podcasts/${podcast.id}`}
            >
              <img
                className="lg:h-56 lg:w-56 h-32 w-32 flex-none object-cover object-center rounded"
                src={`http://localhost:8080/img/${podcast.id}p-500x500.jpg`}
              />
            </Link>
          </div>
        ))}
      </div>
    )
  }
}
