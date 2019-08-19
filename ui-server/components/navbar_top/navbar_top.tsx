import React, { Component } from 'react';
import SearchBarMobile from './components/search_bar_mobile'
import AppHeader from './components/app_header';

interface Props {
  showFullSearchBar: boolean;
  toggleFullSearchBar: () => void;
}

interface State {
  searchText: string;
}

export default class TopNavbar extends Component<Props, State> {
  state = {
    searchText: ""
  }

  handleSearchTextChange = (e: React.FormEvent<HTMLInputElement>): void => {
    e.preventDefault()
    this.setState({
      searchText: e.currentTarget.value
    })
  }

  handleSearchTextSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault()
    console.log(this.state.searchText)
  }
  
  render() {
    const { searchText } = this.state

    return <header className="fixed top-0 left-0 h-12 w-full bg-white lg:border-none">
      { false && this.props.showFullSearchBar
        ? <SearchBarMobile 
            searchText={searchText}
            handleSearchTextChange={this.handleSearchTextChange}
            handleSearchTextSubmit={this.handleSearchTextSubmit}
          />
        : <AppHeader
            searchText={searchText}
            handleSearchTextChange={this.handleSearchTextChange}
            handleSearchTextSubmit={this.handleSearchTextSubmit}
          />
      }
    </header>
  }
}
