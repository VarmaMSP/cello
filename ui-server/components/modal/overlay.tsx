import { Component } from 'react'

interface Props {
  children: JSX.Element
  background: string
}

export default class extends Component<Props> {
  componentDidMount() {
    document.body.style['overflow'] = 'hidden'
  }

  componentWillUnmount() {
    document.body.style['overflow'] = 'auto'
  }

  render() {
    const { children, background } = this.props
    return (
      <>
        <div className="fixed inset-0 bg-white z-10" style={{ background }} />
        {children}
      </>
    )
  }
}
