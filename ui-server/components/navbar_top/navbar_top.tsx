import ButtonWithIcon from 'components/button_with_icon'
import SignInButton from 'components/sign_in_button'
import Router from 'next/router'
import React, { Component } from 'react'
import AppLogo from './components/app_logo'
import SearchBar from './components/search_bar'
import FullWidthSearchBar from './components/search_bar_full_width'
import UserSettings from './components/user_settings'

export interface StateToProps {
  userSignedIn: boolean
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
    const { searchText, userSignedIn } = this.props
    const { showFullWidthSearchBar } = this.state

    if (showFullWidthSearchBar) {
      return (
        <header className="fixed top-0 left-0 h-12 w-full bg-white">
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
      <header className="fixed top-0 left-0 flex justify-between items-center w-full lg:h-14 h-12 lg:pl-56 lg:pr-5 md:px-10 px-4 bg-white">
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
        <div className="md:w-24 w-20 h-8">
          {userSignedIn ? <UserSettings /> : <SignInButton />}
        </div>
      </header>
    )
  }
}
