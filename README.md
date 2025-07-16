# FutureCipherFlow - Secure and Scalable Cryptographic Workflow Engine

FutureCipherFlow is a Go-based project designed to provide a robust and scalable engine for orchestrating complex cryptographic workflows. It aims to simplify the integration and management of various cryptographic operations, allowing developers to easily build secure and efficient applications requiring advanced cryptographic capabilities. This engine allows for defining workflows as directed acyclic graphs (DAGs), where each node represents a specific cryptographic operation. It offers features such as parallel execution, error handling, and detailed logging, enabling reliable and auditable cryptographic processes.

This project is built with security and performance in mind. The use of Go ensures efficient memory management and concurrency, making it suitable for high-throughput cryptographic operations. FutureCipherFlow provides a flexible framework that can be extended to support a wide range of cryptographic algorithms and key management systems. It aims to abstract away the complexities of managing cryptographic keys, algorithms, and protocols, allowing developers to focus on their application logic. By providing a centralized and well-defined workflow engine, FutureCipherFlow enhances the security and maintainability of applications that rely on cryptographic operations.

FutureCipherFlow is intended for use in a variety of applications, including secure data storage, secure communication systems, digital signature management, and privacy-preserving computation. It is designed to be modular and extensible, allowing developers to customize the engine to meet their specific needs. The engine provides a clear separation of concerns between the workflow definition, the cryptographic operations, and the underlying infrastructure, which promotes code reusability and reduces the risk of errors. The project also includes comprehensive documentation and examples to help developers get started quickly.

## Key Features

*   **Workflow Definition as DAGs:** Defines cryptographic workflows as directed acyclic graphs (DAGs), allowing for complex and parallel operations. Each node in the graph represents a specific cryptographic task. The DAG structure ensures that there are no circular dependencies, which prevents infinite loops and simplifies the workflow management.
*   **Parallel Execution:** Leverages Go's concurrency features to execute independent cryptographic operations in parallel, improving performance and reducing processing time. The engine intelligently manages the execution of tasks based on their dependencies, ensuring that tasks are executed in the correct order.
*   **Error Handling and Recovery:** Implements robust error handling and recovery mechanisms to ensure the reliability of cryptographic workflows. The engine includes features such as retry policies, dead-letter queues, and circuit breakers to handle transient errors and prevent cascading failures.
*   **Detailed Logging and Auditing:** Provides comprehensive logging and auditing capabilities to track the execution of cryptographic workflows and identify potential security issues. The engine logs all significant events, including the start and end of tasks, error messages, and performance metrics.
*   **Extensible Architecture:** Supports a wide range of cryptographic algorithms and key management systems through a modular and extensible architecture. Developers can easily add new algorithms and key management providers without modifying the core engine.
*   **Key Management Integration:** Integrates with various key management systems (e.g., HashiCorp Vault, AWS KMS) to securely manage cryptographic keys. The engine provides a consistent interface for accessing keys, regardless of the underlying key management system.
*   **API for Workflow Management:** Exposes a well-defined API for managing cryptographic workflows, including creating, starting, stopping, and monitoring workflows. The API allows developers to integrate the engine into their existing applications and infrastructure.

## Technology Stack

*   **Go:** The core programming language used for building the engine due to its performance, concurrency features, and strong standard library. Go's efficient memory management and garbage collection make it well-suited for high-throughput cryptographic operations.
*   **gRPC:** Used for defining the API and communication between different components of the system. gRPC provides a high-performance and type-safe mechanism for communication, which is essential for building reliable and scalable applications.
*   **Protocol Buffers:** Used for defining the data structures and messages exchanged between components. Protocol Buffers provide a compact and efficient serialization format, which reduces network overhead and improves performance.
*   **etcd/Consul (Optional):** Used for service discovery and configuration management. These systems allow the engine to dynamically discover and configure its components, which improves scalability and resilience.
*   **logrus:** Used for logging, providing structured and configurable logging capabilities. Logrus allows developers to easily customize the logging output and integrate it with various logging systems.

## Installation

1.  **Prerequisites:** Ensure you have Go (version 1.18 or later) installed and configured on your system. Verify by running `go version`. Also, ensure that you have `git` installed.

2.  **Clone the Repository:** Clone the FutureCipherFlow repository from GitHub using the following command:
   git clone https://github.com/jjfhwang/FutureCipherFlow.git

3.  **Navigate to the Project Directory:** Change your current directory to the cloned repository:
   cd FutureCipherFlow

4.  **Download Dependencies:** Use Go modules to download the required dependencies:
   go mod download

5.  **Build the Project:** Build the FutureCipherFlow executable using the following command:
   go build -o futurecipherflow .

## Configuration

FutureCipherFlow can be configured using environment variables. The following environment variables are supported:

*   `FCF_WORKFLOW_DIR`: Specifies the directory where workflow definitions are stored (default: `./workflows`).
*   `FCF_LOG_LEVEL`: Sets the logging level (e.g., `debug`, `info`, `warn`, `error`, `fatal`) (default: `info`).
*   `FCF_GRPC_PORT`: Specifies the port on which the gRPC server listens (default: `50051`).
*   `FCF_KEY_MANAGEMENT_SYSTEM`: Specifies the key management system to use (e.g., `vault`, `aws_kms`) (default: `vault`).
*   `FCF_VAULT_ADDR` (Optional): Specifies the address of the HashiCorp Vault server (required if `FCF_KEY_MANAGEMENT_SYSTEM` is set to `vault`).
*   `FCF_AWS_REGION` (Optional): Specifies the AWS region (required if `FCF_KEY_MANAGEMENT_SYSTEM` is set to `aws_kms`).

Example:
  export FCF_LOG_LEVEL=debug
  export FCF_GRPC_PORT=60000

## Usage

1.  **Define a Workflow:** Create a workflow definition file in the `workflows` directory. The workflow definition should be in JSON or YAML format and specify the nodes and edges of the DAG. Example workflow definition (example.json):

{
    "name": "encryption_workflow",
    "nodes": [
        {"id": "generate_key", "type": "generate_key", "algorithm": "AES256"},
        {"id": "encrypt_data", "type": "encrypt", "input": "data.txt", "key_id": "generate_key"}
    ],
    "edges": [
        {"source": "generate_key", "target": "encrypt_data"}
    ]
}

2.  **Start the Engine:** Run the compiled executable:
   ./futurecipherflow

3.  **Interact with the API:** Use the gRPC API to manage workflows. You can use tools like `grpcurl` or generate client code using `protoc`.

Example using `grpcurl`:

grpcurl -plaintext -d '{"workflow_name": "encryption_workflow"}' localhost:50051 FutureCipherFlow.WorkflowService/StartWorkflow

(Remember to replace 50051 with the configured FCF_GRPC_PORT if different). The exact API calls and parameters will depend on the specific gRPC service definition, which would be defined in a .proto file within the project. This file should be consulted for specific message structures.

## Contributing

We welcome contributions to FutureCipherFlow! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes and write tests.
4.  Ensure that all tests pass.
5.  Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/FutureCipherFlow/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for providing the tools and libraries that made this project possible.