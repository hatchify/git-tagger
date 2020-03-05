# Git Tagger
Git tagger helps with incrementing git tags in SEMVER format

## Usage
### Increment patch
```bash
# Patch is the default type, so it's not needed to be explicitly set
git-tagger

# If you want to be verbose with your version type
git-tagger --type patch
```

### Increment minor
```bash
# Set type as "minor"
git-tagger --type minor
```

### Increment major
```bash
# Set type as "major"
git-tagger --type major
```
