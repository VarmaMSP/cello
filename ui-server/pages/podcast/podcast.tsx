import React, { Component } from 'react'

interface Props {
  getPodcast: (id: string) => void
}

export default class PodcastPage extends Component<Props, {}> {
  componentDidMount() {
    this.props.getPodcast('blj6eu3c1osnagmbgu10')
  }

  render() {
    return <div>Im Mr MeeSeeks Look at me</div>
  }
}
