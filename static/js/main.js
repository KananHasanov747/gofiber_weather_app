document.addEventListener("alpine:init", () => {
  Alpine.data("autocomplete", () => ({
    version: "1",
    searchAPI: "search",
    weatherAPI: "weather",
    expanded: false,
    query: "",
    locations: [],
    highlightedIndex: 0,
    index: null,
    get fetchLocations() {
      if (this.query.length > 0) {
        fetch(`api/v${this.version}/${this.searchAPI}/${this.query}`)
          .then((response) => response.text())
          .then((list) => {
            this.locations = JSON.parse(list);
            this.highlightedIndex = 0;
          });
      } else {
        this.locations = [];
      }
    },
    get showSuggestions() {
      if (this.tmpLocations) this.locations = this.tmpLocations;
    },
    get hideSuggestions() {
      this.tmpLocations = this.locations;
      this.locations = [];
    },
    get highlightNext() {
      if (this.highlightedIndex < this.locations.length - 1) {
        this.highlightedIndex++;
      } else {
        this.highlightedIndex = 0;
      }
      //this.$refs.suggestionsContainer.children[
      //  this.highlightedIndex
      //].scrollIntoView({ block: "nearest" });
    },
    get highlightPrevious() {
      if (this.highlightedIndex > 0) {
        this.highlightedIndex--;
      } else {
        this.highlightedIndex = this.locations.length - 1;
      }
      //this.$refs.suggestionsContainer.children[
      //  this.highlightedIndex
      //].scrollIntoView({ block: "nearest" });
    },
    // Send the response to the backend and retrieve an updated data
    get selectLocation() {
      const defaultSelected =
        this.locations[this.index ?? this.highlightedIndex];
      const selectedCity = defaultSelected.name;
      const selectedCountry = defaultSelected.country;
      const selectedLatitude = defaultSelected.latitude;
      const selectedLongitude = defaultSelected.longitude;

      if (selectedCity && selectedCountry) {
        // Use cookies to set the values unchanged after reload
        let expires = new Date();
        expires.setTime(expires.getTime() + 7 * 86400 * 1000); // 7 days * 24 hours * 1 sec
        document.cookie =
          "city=" + selectedCity + "; expires=" + expires.toUTCString();
        document.cookie =
          "country=" + selectedCountry + "; expires=" + expires.toUTCString();
        document.cookie =
          "lat=" + selectedLatitude + "; expires=" + expires.toUTCString();
        document.cookie =
          "lon=" + selectedLongitude + "; expires=" + expires.toUTCString();

        window.history.pushState({}, "", this.url);
        htmx.ajax(
          "GET",
          `api/v${this.version}/${this.weatherAPI}/${selectedCity}/${selectedCountry}/${selectedLatitude}/${selectedLongitude}`,
          {
            handler: (_, response) => {
              this.data = JSON.parse(response.xhr.responseText);
            },
          },
        );
        this.query = "";
        this.locations = [];
      }
    },
  }));

  Alpine.data("tabs", () => ({
    tabSelected: 0,
    prevTabSelected: 0,
    tabSectionClicked(index) {
      this.prevTabSelected = this.tabSelected;
      this.tabSelected = index;
      const tabSection = this.$refs.tabSections.children[index + 1]; // the first element is <template> itself
      this.tabRepositionMarker(tabSection);
    },
    tabRepositionMarker(tabSection) {
      this.$refs.tabMarker.style.width = tabSection.offsetWidth + "px";
      this.$refs.tabMarker.style.left = tabSection.offsetLeft + "px";
    },
    tabContentActive(tabContent) {
      return (
        parseInt(this.tabSelected) ===
        parseInt(tabContent.id.replace(this.tabId + "-content-", ""), 10)
      );
    },
  }));
});
