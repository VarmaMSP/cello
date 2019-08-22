import React, { Component } from 'react';
import AppLogo from './components/app_logo'
import FullWidthSearchBar from './components/search_bar_full_width';
import SearchBar from './components/search_bar'
import SignInButton from './components/sign_in_button';
import ButtonWithIcon from '../button_with_icon';

interface Props {
}

interface State {
  searchText: string;
  showFullWidthSearchBar: boolean;
}

export default class TopNavbar extends Component<Props, State> {
  state = {
    searchText: "",
    showFullWidthSearchBar: false
  }

  handleSearchBarCollapse = () => {
    const { showFullWidthSearchBar } = this.state
    this.setState({showFullWidthSearchBar: !showFullWidthSearchBar})
  }

  handleSearchTextChange = (e: React.FormEvent<HTMLInputElement>) => {
    e.preventDefault()
    this.setState({searchText: e.currentTarget.value})
  }

  handleSearchTextSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    console.log("subbmited: ", this.state.searchText)
  }
  
  render() {
    const { showFullWidthSearchBar, searchText } = this.state

    if (showFullWidthSearchBar) {
      return <header className="fixed top-0 left-0 h-12 w-full bg-white">
        <FullWidthSearchBar
          searchText={searchText}
          handleCollapse={this.handleSearchBarCollapse}
          handleSearchTextChange={this.handleSearchTextChange}
          handleSearchTextSubmit={this.handleSearchTextSubmit}
        />
      </header>
    }

    return <header className="fixed top-0 left-0 flex justify-between items-center w-full h-12 lg:pl-56 lg:pr-4 md:px-10 px-3 bg-white">
      <div className="lg:hidden w-20">
        <ButtonWithIcon className="w-5" icon="search" onClick={this.handleSearchBarCollapse}/>
      </div>
      <div className="lg:hidden">
        <AppLogo/>
      </div>
      <div className="lg:block hidden mx-3 my-1">
        <SearchBar
          searchText={searchText}
          handleSearchTextChange={this.handleSearchTextChange}
          handleSearchTextSubmit={this.handleSearchTextSubmit}
        />
      </div>
      <SignInButton/>
    </header>
  }
}
