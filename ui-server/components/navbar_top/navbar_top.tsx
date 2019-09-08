import React, { Component } from 'react'
import AppLogo from './components/app_logo'
import FullWidthSearchBar from './components/search_bar_full_width'
import SearchBar from './components/search_bar'
import SignInButton from './components/sign_in_button'
import ButtonWithIcon from '../button_with_icon'
import Router from 'next/router'

export interface StateToProps {
  searchText: string
}

export interface DispatchToProps {
  searchTextChange: (text: string) => void
}

interface Props extends StateToProps, DispatchToProps {}

interface State {
  showFullWidthSearchBar: boolean
}

export default class TopNavbar extends Component<Props, State> {
  state = {
    showFullWidthSearchBar: false,
  }

  handleSearchBarCollapse = () => {
    const { showFullWidthSearchBar } = this.state
    this.setState({ showFullWidthSearchBar: !showFullWidthSearchBar })
  }

  handleSearchTextChange = (e: React.FormEvent<HTMLInputElement>) => {
    e.preventDefault()
    this.props.searchTextChange(e.currentTarget.value)
  }

  handleSearchTextSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const { showFullWidthSearchBar } = this.state
    if (showFullWidthSearchBar) {
      this.setState({ showFullWidthSearchBar: false })
    }

    const searchText = this.props.searchText.trim()
    if (searchText.length > 0) {
      Router.push({
        pathname: '/results',
        query: { search_query: searchText },
      })
    }
  }

  render() {
    const { searchText } = this.props
    const { showFullWidthSearchBar } = this.state

    if (showFullWidthSearchBar) {
      return (
        <header className="fixed top-0 left-0 h-12 w-full bg-white z-50">
          <FullWidthSearchBar
            searchText={searchText}
            handleCollapse={this.handleSearchBarCollapse}
            handleSearchTextChange={this.handleSearchTextChange}
            handleSearchTextSubmit={this.handleSearchTextSubmit}
          />
        </header>
      )
    }

    return (
      <header className="fixed top-0 left-0 flex justify-between items-center w-full lg:h-14 h-12 lg:pl-56 lg:pr-5 md:px-10 px-4 bg-white z-50">
        <div className="lg:hidden w-20">
          <ButtonWithIcon
            className="w-6"
            icon="search"
            onClick={this.handleSearchBarCollapse}
          />
        </div>
        <div className="lg:hidden">
          <AppLogo />
        </div>
        <div className="lg:block hidden mx-3 ">
          <SearchBar
            searchText={searchText}
            handleSearchTextChange={this.handleSearchTextChange}
            handleSearchTextSubmit={this.handleSearchTextSubmit}
          />
        </div>
        <SignInButton />
      </header>
    )
  }
}
