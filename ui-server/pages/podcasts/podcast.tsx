import React, { Component } from 'react'
import Link from 'next/link'
import { PageContext } from 'types/next'

interface Props {
  id: string
  getPodcast: (id: string) => void
  isServer: boolean
}

export default class PodcastPage extends Component<Props, {}> {
  static async getInitialProps({
    query,
    isServer,
  }: PageContext): Promise<Partial<Props>> {
    const podcastId = query['id'] as string
    return { id: podcastId, isServer }
  }

  componentDidUpdate(prevProps: Props) {
    if (this.props.id != prevProps.id) {
      this.props.getPodcast(this.props.id)
    }
  }

  componentDidMount() {
    this.props.getPodcast(this.props.id)
  }

  render() {
    return (
      <>
        <Link
          href={{
            pathname: '/podcasts',
            query: { id: 'blj7283c1osnagmfk210' },
          }}
          as={`/podcasts/${'blj7283c1osnagmfk210'}`}
        >
          <a>Im Mr MeeSeeks Look at me</a>
        </Link>
        <div>Im Mr MeeSeeks Look at me</div>
        <div>{`Im a at ${this.props.id}`}</div>
      </>
    )
  }
}
