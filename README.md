# URL Shortener Service

The URL Shortener Service is a system that generates shorter and unique aliases for long URLs, allowing users to access the original links through these shortened URLs. The service provides functionality for generating custom short links, specifying expiration times for links, and tracking analytics such as the number of times a redirection occurs. Additionally, the service is accessible through REST APIs, enabling integration with other services.

## Functional Requirements

1. Given a URL, the service should generate a shorter and unique alias, known as a short link.
2. When users access a short link, the service should redirect them to the original long URL.
3. Users should have the option to choose a custom short link for their URL.
4. Links should have an expiration time, which can be specified by the users.

## Non-Functional Requirements

1. **High Availability:** The system should be highly available to ensure uninterrupted URL redirection even during periods of high traffic or failures.
2. **Real-Time Redirection:** URL redirection should occur with minimal latency, providing a seamless user experience.
3. **Unpredictable Short Links:** Shortened links should not be guessable or predictable to prevent unauthorized access or abuse.

## Extended Requirements

1. **Analytics:** The system should track analytics, such as the number of times a redirection has occurred, to provide insights into link usage and performance.
2. **REST APIs:** The service should expose REST APIs, allowing other services to interact with it programmatically.

## Usage

To use the URL Shortener Service, you can interact with the provided REST APIs. Detailed documentation for the APIs can be found in the [API Documentation](api.md) file.

## Development

### Prerequisites

- Go (version X.X.X)
- Database (e.g., PostgreSQL)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/url-shortener.git
