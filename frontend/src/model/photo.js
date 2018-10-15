import Abstract from 'model/abstract';

class Photo extends Abstract {
    getEntityName() {
        return this.PhotoTitle;
    }

    getId() {
        return this.ID;
    }

    getGoogleMapsLink() {
        return 'https://www.google.com/maps/place/' + this.PhotoLat + ',' + this.PhotoLong;
    }

    getThumbnailUrl(type, size) {
        return '/api/v1/thumbnails/' + type + '/' + size + '/' + this.FileHash;
    }

    getThumbnailSrcset() {
        const result = [];

        result.push(this.getThumbnailUrl('fit', 320) + ' 320w');
        result.push(this.getThumbnailUrl('fit', 500) + ' 500w');
        result.push(this.getThumbnailUrl('fit', 720) + ' 720w');
        result.push(this.getThumbnailUrl('fit', 1280) + ' 1280w');
        result.push(this.getThumbnailUrl('fit', 1920) + ' 1920w');
        result.push(this.getThumbnailUrl('fit', 2560) + ' 2560w');
        result.push(this.getThumbnailUrl('fit', 3840) + ' 3840w');

        return result.join(', ');
    }

    getThumbnailSizes() {
        const result = [];

        result.push('(min-width: 2560px) 3840px');
        result.push('(min-width: 1920px) 2560px');
        result.push('(min-width: 1280px) 1920px');
        result.push('(min-width: 720px) 1280px');
        result.push('(min-width: 500px) 720px');
        result.push('(min-width: 320px) 500px');
        result.push('320px');

        return result.join(', ');
    }

    getLocation() {
        const location = [];

        if (this.LocationID) {
            if (this.LocName && !this.LocCity && !this.LocCounty) {
                location.push(this.LocName)
            } else if (this.LocCity) {
                location.push(this.LocCity)
            } else if (this.LocCounty) {
                location.push(this.LocCounty)
            }

            if (this.LocState && this.LocState !== this.LocCity) {
                location.push(this.LocState)
            }

            if (this.LocCountry) {
                location.push(this.LocCountry)
            }
        } else if (this.CountryName) {
            location.push(this.CountryName)
        } else {
            location.push('Unknown')
        }

        return location.join(', ');
    }

    getFullLocation() {
        const location = [];

        if (this.LocationID) {
            if (this.LocName) {
                location.push(this.LocName)
            }

            if (this.LocCity) {
                location.push(this.LocCity)
            }

            if (this.LocPostcode) {
                location.push(this.LocPostcode)
            }

            if (this.LocCounty) {
                location.push(this.LocCounty)
            }

            if (this.LocState) {
                location.push(this.LocState)
            }

            if (this.LocCountry) {
                location.push(this.LocCountry)
            }
        } else if (this.CountryName) {
            location.push(this.CountryName)
        } else {
            location.push('Unknown')
        }

        return location.join(', ');
    }

    getCamera() {
        if (this.CameraModel) {
            return this.CameraModel
        }

        return 'Unknown'
    }

    static getCollectionResource() {
        return 'photos';
    }

    static getModelName() {
        return 'Photo';
    }
}

export default Photo;
